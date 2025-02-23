# Use Go base image
FROM golang:1.20-alpine

# Install curl and other necessary dependencies
RUN apk add --no-cache curl

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Expose the application port
EXPOSE 8080

# Build the Go application
RUN go build -o app ./cmd/main.go

# Run the application
CMD ["./app"]
