FROM golang:alpine

WORKDIR /opt

COPY go.mod go.sum main.go ./

# Download all dependencies
RUN go mod download

# Build the Go application
RUN go build

# Command to run the executable
CMD ["./GotifyDiscordRedirector"]