FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o weather-cli ./cmd/weather-cli

FROM alpine:latest

COPY --from=builder /app/weather-cli /usr/local/bin/weather-cli

ENTRYPOINT ["weather-cli"]
