# FROM golang:1.18.0-bullseye AS builder
FROM golang:1.18.0-bullseye

WORKDIR /go/src/github.com/moromin/PFC-balancer

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/user ./services/user

# CMD ["/usr/bin/user", "-p", "4000", "-dbAddr", "db:50000"]
ENTRYPOINT ["/usr/bin/user"]
