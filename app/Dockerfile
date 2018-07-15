FROM golang:latest

ENV GOPATH $GOPATH:/go/src
RUN apt-get update && \
    apt-get upgrade -y

RUN mkdir /go/src/app

EXPOSE 5000
