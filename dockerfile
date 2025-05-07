# Dockerfile
FROM golang:1.24.2-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o investor-api .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/investor-api .

COPY .env.example .env

EXPOSE 8080

CMD ["./investor-api"]
