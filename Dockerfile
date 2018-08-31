FROM golang:1.8-alpine

RUN apk update && apk add bash

WORKDIR /go/src/github.com/gregmankes/solver