# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN go build -o app

# Stage 2: Create a minimal production image
FROM alpine:latest
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .
COPY ./config.json /app/config.json

# Expose the port (if your application listens on a port)
EXPOSE 8080

# Run the Go binary
CMD ["./app"]
