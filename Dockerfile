FROM golang:1.16.4-alpine

RUN apk update

RUN apk upgrade

COPY . ./project/golang-logstash