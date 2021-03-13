FROM golang:1.13.5-alpine

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/src/app

COPY ./main.go .

RUN apk add git \
    && go get . \
    && go build

FROM alpine:3.10.3

WORKDIR /app

COPY --from=0 /go/bin/app /app

EXPOSE 8000/tcp

ENTRYPOINT ./app