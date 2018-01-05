FROM golang:latest

ENV GOPATH $GOPATH:/go/src
RUN apt-get update && \
    apt-get upgrade -y

#install library
RUN go get github.com/go-kit/kit/endpoint
RUN go get github.com/go-kit/kit/transport/http
RUN go get github.com/go-kit/kit/log
RUN go get github.com/go-kit/kit/metrics
RUN go get github.com/go-kit/kit/metrics/prometheus
RUN go get github.com/prometheus/client_golang/prometheus

RUN mkdir /go/src/app

EXPOSE 5000
