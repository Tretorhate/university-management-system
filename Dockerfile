# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o university-management-system ./cmd/api

# Stage 2: Run the application with a smaller image
FROM alpine:latest

WORKDIR /root/

# Copy the binary and .env file from the builder
COPY --from=builder /app/university-management-system .
COPY --from=builder /app/.env .

# Expose the app port
EXPOSE 8080

# Start the app
CMD ["./university-management-system"]
