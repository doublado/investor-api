# Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Required for go-openai and godotenv
RUN apk add --no-cache git

# Copy and initialize Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go binary
RUN go build -o investor-api .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy binary and files
COPY --from=builder /app/investor-api .

# Optional: copy static/config files if needed
COPY .env.example .env

EXPOSE 8080

CMD ["./investor-api"]
