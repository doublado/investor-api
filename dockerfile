# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Disable CGO for static binary compatible with distroless
ENV CGO_ENABLED=0
RUN go build -o app

# Minimal runtime stage using distroless
FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/app/app"]
