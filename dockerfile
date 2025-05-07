# Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o investor-api .

EXPOSE 8080

CMD ["./investor-api"]
