
FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd

FROM alpine:latest

COPY --from=builder /app/cmd/main /app/cmd/main

WORKDIR /app

CMD ["./cmd/main"]
