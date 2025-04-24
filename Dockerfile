FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o university-management-system ./cmd/api

# Use a small alpine image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/university-management-system .
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./university-management-system"]