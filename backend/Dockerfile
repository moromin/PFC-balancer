FROM golang:1.17-alpine as dev

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/cespare/reflex
ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY ./backend/go.* ./backend/main.go ./

RUN go mod download

EXPOSE 8000


FROM golang:1.17-alpine as builder

ENV ROOT=go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY ./backend/ ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux fo build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8000
CMD ["/go/src/app/binary"]
