FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o weather-consumer .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/weather-consumer .
EXPOSE 9090
CMD ["./weather-consumer"]