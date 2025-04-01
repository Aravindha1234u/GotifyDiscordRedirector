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
