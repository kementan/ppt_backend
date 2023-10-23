FROM golang:1.21-alpine3.18

RUN apk update && apk add --no-cache git build-base

ENV CGO_ENABLED 1

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ppt_backend_golang ./cmd/

EXPOSE 80

CMD ["./ppt_backend_golang"]
