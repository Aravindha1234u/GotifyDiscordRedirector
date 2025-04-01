package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// Message represents the structure of the message received from Gotify
type Message struct {
	ID       int       `json:"id"`
	AppID    int       `json:"appid"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	Priority int       `json:"priority"`
	Date     time.Time `json:"date"`
}

type DiscordMessage struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds,omitempty"`
}
type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}

func main() {
	// Create a channel to listen for OS signals (like Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var gotifyURL string
	var discordWebhookURL string

	// Use command line arguments or environment variables
	flag.StringVar(&gotifyURL, "gotify", os.Getenv("GOTIFY_WS_URL"), "Gotify WebSocket URL. Example: wss://gotify.example.com/stream?token=abcdefghijklmnop")
	flag.StringVar(&discordWebhookURL, "discord", os.Getenv("DISCORD_WEBHOOK_URL"), "Discord Webhook URL. Example: https://discord.com/api/webhooks/123456789012345678/abcdefghijklmnopqrstuvwxyz")
	flag.Parse()

	if gotifyURL == "" || discordWebhookURL == "" {
		log.Fatal("Gotify URL and Discord Webhook URL must be provided (environment or command line).")
	}

	log.Printf("connecting to %s", gotifyURL)
	c, _, err := websocket.DefaultDialer.Dial(gotifyURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Start a goroutine to read messages from the WebSocket
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Fatal("Error reading message:", err)
				os.Exit(1) // Exit if there's an error reading the message
			}

			// Parse the JSON message and send it to Discord
			var gotifyMsg Message
			if err := json.Unmarshal(message, &gotifyMsg); err != nil {
				log.Fatal("Error parsing message:", err)
				continue
			}

			// Create a Discord message
			discordMsg := DiscordMessage{}

			// Create a Discord embed based on priority
			switch {
			case gotifyMsg.Priority == 0:
				discordMsg.Embeds = []Embed{{
					Title:       gotifyMsg.Title,
					Description: gotifyMsg.Message,
					Color:       0x808080, // Gray
				}}
			case gotifyMsg.Priority >= 1 && gotifyMsg.Priority <= 3:
				discordMsg.Embeds = []Embed{{
					Title:       fmt.Sprintf("ℹ️ %s", gotifyMsg.Title),
					Description: gotifyMsg.Message,
					Color:       0x00BFFF, // Deep sky blue
				}}
			case gotifyMsg.Priority >= 4 && gotifyMsg.Priority <= 7:
				discordMsg.Embeds = []Embed{{
					Title:       fmt.Sprintf("🔔 %s", gotifyMsg.Title),
					Description: gotifyMsg.Message,
					Color:       0xFFA500, // Orange
				}}
			case gotifyMsg.Priority >= 8 && gotifyMsg.Priority <= 10:
				discordMsg.Embeds = []Embed{{
					Title:       fmt.Sprintf("🚨 %s", gotifyMsg.Title),
					Description: gotifyMsg.Message,
					Color:       0xFF0000, // Red
				}}
			default:
				discordMsg.Content = fmt.Sprintf("**%s**\n\n%s", gotifyMsg.Title, gotifyMsg.Message)
			}

			discordMsgJSON, err := json.Marshal(discordMsg)
			if err != nil {
				log.Fatal("Error marshalling Discord message: %v", err)
				continue
			}

			resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(discordMsgJSON))
			if err != nil {
				log.Fatal("Error sending message to Discord: %v", err)
			}
			defer resp.Body.Close()

		}
	}()

	<-interrupt // Block until an interrupt signal is received

	// Cleanly close the connection by sending a close message.
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Fatal("write close:", err)
		return
	}
	<-done //Wait for the read goroutine to finish.
}
