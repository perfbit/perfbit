# Dockerfile

# Build stage
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Install git (required for go mod download)
RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/auth-service .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./auth-service"]