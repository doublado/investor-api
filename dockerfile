# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Disable CGO for static binary
ENV CGO_ENABLED=0
RUN go build -o app

# Runtime stage
FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/app/app"]
