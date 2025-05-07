# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# Final stage
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080

CMD ["./app"]
