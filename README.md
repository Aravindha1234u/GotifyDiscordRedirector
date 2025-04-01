# Gotify to Discord Forwarder
This Go program forwards messages from a Gotify WebSocket stream to a Discord webhook, providing customizable formatting based on Gotify message priority.

## Motivation
This tool was created to enhance the readability and accessibility of notifications from a Gotify push server. While Gotify is excellent for centralizing push notifications, its native display on PCs and phones can sometimes be less than ideal. Discord, on the other hand, is frequently running in the background on both desktops and mobile devices, offering a more consistent and user-friendly notification experience.

By redirecting Gotify messages to a dedicated Discord channel, you can:

* **Centralize Notifications:** Keep all your important alerts in one place, alongside your other Discord communications.
* **Ensure Active Monitoring:** Leverage Discord's always-on nature to ensure you never miss a critical notification, even when your Gotify client might not be actively in focus.
* **Customize Alerts:** Format the alerts based on priority to visually highlight the importance of each notification.
* **Improve Readability:** Discord's rich text formatting and embed capabilities make it easier to digest notification content.

## Features
* **Priority-Based Formatting:** Formats Discord messages based on Gotify priority levels, including colors and emojis.
* **Docker Support:** Includes a Dockerfile for easy containerization.
* **Detailed Logging:** Uses Go's `log` package for clear and informative logging.
* **Environment Variables or Command-Line Arguments:** Configurable via environment variables or command-line flags.

## Gotify Priority Level Handling
The program formats Discord messages based on the following Gotify priority levels:

* **0 (Silent):** Gray embed.
* **1-3 (Low):** Deep sky blue embed with an "ℹ️" emoji.
* **4-7 (Default/Normal):** Orange embed with a "🔔" emoji.
* **8-10 (High):** Red embed with an "🚨" emoji.
* **Other:** If the priority is outside the 0-10 range, a standard Discord message is sent with the priority number.

## Configuration
You can configure the Gotify and Discord URLs using either environment variables or command-line arguments.

### Environment Variables
* `GOTIFY_WS_URL`: The Gotify WebSocket URL (e.g., `wss://your-gotify-url/stream?token=your-token`).
* `DISCORD_WEBHOOK_URL`: Your Discord webhook URL (e.g., `https://discord.com/api/webhooks/your-webhook-id/your-webhook-token`).

### Command-Line Arguments
* `-gotify`: The Gotify WebSocket URL.
* `-discord`: Your Discord webhook URL.

## Installation

### Binary Release
1.  Navigate to the [GitHub Releases](https://github.com/Aravindha1234u/GotifyDiscordRedirector/releases) page.
2.  Download the appropriate binary for your operating system and architecture.
3.  Make the binary executable (if necessary): `chmod +x GotifyDiscordRedirector`.
4.  Run the binary: `./GotifyDiscordRedirector -gotify="ws://your-gotify-url/stream?token=your-token" -discord="https://discord.com/api/webhooks/your-webhook-id/your-webhook-token"`.

### Linux Service (systemd)
To run the `GotifyDiscordRedirector` binary as a background service on Linux using `systemd`, follow these steps:

1.  **Create a Service User (Recommended):**
```bash
sudo adduser --system --group --no-create-home gotify-discord
```

2.  **Copy the Binary:**
Copy the `GotifyDiscordRedirector` binary to a suitable location, such as `/opt/`:

```bash
sudo cp GotifyDiscordRedirector /opt/
sudo chown gotify-discord:gotify-discord /opt/GotifyDiscordRedirector
```

3.  **Create the systemd Service File:**
Create a service file at `/etc/systemd/system/gotify-discord.service`:

```ini
[Unit]
Description=Gotify to Discord Forwarder
After=network.target

[Service]
User=gotify-discord
Group=gotify-discord
ExecStart=/opt/GotifyDiscordRedirector -gotify="ws://your-gotify-url/stream?token=your-token" -discord="https://discord.com/api/webhooks/your-webhook-id/your-webhook-token"
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

**Important:** Replace the appropriate command-line flags or ensure the environment variables are set in the systemd environment.

4.  **Reload systemd and Enable the Service:**
```bash
sudo systemctl daemon-reload
sudo systemctl enable gotify-discord.service
```

5.  **Start the Service:**
```bash
sudo systemctl start gotify-discord.service
```

6.  **Check the Service Status:**
```bash
sudo systemctl status gotify-discord.service
```

7.  **View Logs:**
```bash
sudo journalctl -u gotify-discord.service
```

### Docker Image
1.  Pull the Docker image from the Container Registry (GCR): `docker pull aravindha1234u/gotifydiscordredirector:latest`.
2.  Run the Docker container, providing the necessary environment variables:
```bash
sudo docker run -e GOTIFY_WS_URL="ws://your-gotify-url/stream?token=your-token" -e DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/your-webhook-id/your-webhook-token" aravindha1234u/gotifydiscordredirector:latest
```

## Build

### Prerequisites
* Go (1.18 or later)
* Docker (for containerization)
* Gotify server
* Discord webhook URL

### Build from Source
1.  Clone the repository:
```bash
git clone https://github.com/Aravindha1234u/GotifyDiscordRedirector.git
cd GotifyDiscordRedirector
```

2.  Install dependencies:
```bash
go mod download
```

3.  Build the application:
```bash
go build
```

4.  Run the application:
```bash
./GotifyDiscordRedirector -gotify="ws://your-gotify-url/stream?token=your-token" -discord="https://discord.com/api/webhooks/your-webhook-id/your-webhook-token"
```

### Build Docker Image

1.  Clone the repository (if you haven't already).
2.  Navigate to the project directory.
3.  Build the Docker image:
```bash
docker build -t GotifyDiscordRedirector .
```

4.  Run the Docker container, providing the necessary environment variables:
```bash
docker run -e GOTIFY_WS_URL="ws://your-gotify-url/stream?token=your-token" -e DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/your-webhook-id/your-webhook-token" GotifyDiscordRedirector
```
