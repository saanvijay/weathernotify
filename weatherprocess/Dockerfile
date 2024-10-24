FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o weather-notify .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/weather-notify .
EXPOSE 8080
CMD ["./weather-notify"]