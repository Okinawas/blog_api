FROM golang:latest

ENV GO111MODULE=on

RUN mkdir /go/src/github.com\
    && mkdir /go/src/github.com/nntakuya\
    && mkdir /go/src/github.com/nntakuya/todo_app\
    && apt-get update -qq \
    && apt-get install -y mariadb-client vim

WORKDIR /go/src/github.com/nntakuya/todo_app

ADD . /go/src/github.com/nntakuya/todo_app


