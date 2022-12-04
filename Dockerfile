# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /simple-ddns-client

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /simple-ddns-client

CMD [ "/simple-ddns-client" ]