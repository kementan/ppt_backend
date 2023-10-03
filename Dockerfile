# Use the official Go image as the base image
FROM golang:1.21-alpine3.18

RUN apk update && apk add --no-cache git build-base

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ppt_backend_binary ./cmd/

EXPOSE 8080

ENTRYPOINT ["./ppt_backend_binary"]
