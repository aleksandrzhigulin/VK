FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main .

EXPOSE 8080 8080
ENTRYPOINT ["/app/main"]