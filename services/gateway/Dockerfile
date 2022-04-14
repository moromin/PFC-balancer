FROM golang:1.18.0-bullseye AS builder

WORKDIR /go/src/github.com/moromin/PFC-balancer

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/gateway ./services/gateway

ENTRYPOINT ["/usr/bin/gateway"]
