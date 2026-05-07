# Build Stage
FROM golang:1.26 AS builder

# Set working directory
WORKDIR /app

# Install Templ CLI
RUN go install github.com/a-h/templ/cmd/templ@latest

# Download Tailwind CSS standalone CLI
# (Adjust version if necessary, v3.4.1 is stable)
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Generate Templ files
RUN templ generate

# Build CSS
# Using the config and input from your project
RUN tailwindcss -i static/css/index.css -o static/css/output.css --minify

# Build the Go application
# CGO_ENABLED=0 for a static binary that works on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder
COPY --from=builder /app/main .

# Copy static assets
COPY --from=builder /app/static ./static

# Copy the initial data file (if it doesn't exist, app might crash)
# Note: In production, you'd usually mount a volume here
COPY --from=builder /app/todos.json .

# Expose the port from main.go
EXPOSE 5000

# Run the application
CMD ["./main"]
