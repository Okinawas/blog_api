FROM golang:latest

ENV GO111MODULE=on

RUN mkdir /go/src/github.com\
    && mkdir /go/src/github.com/Okinawas\
    && mkdir /go/src/github.com/Okinawas/blog_api\
    && apt-get update -qq \
    && apt-get install -y mariadb-client vim

WORKDIR /go/src/github.com/Okinawas/blog_api

ADD . /go/src/github.com/Okinawas/blog_api

