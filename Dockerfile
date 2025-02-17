# Use Golang base image
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o /to-do-app ./cmd/main.go

# Create a smaller image for runtime
FROM alpine:latest
RUN apk update && apk add curl
WORKDIR /root/
COPY --from=builder /to-do-app .

EXPOSE 8080
CMD ["./to-do-app"]

