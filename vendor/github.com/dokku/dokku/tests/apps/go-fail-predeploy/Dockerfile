FROM golang:1.9.1

EXPOSE 3000/tcp

ADD . /go/src/github.com/dokku/apps/web
WORKDIR /go/src/github.com/dokku/apps/web

RUN go install -v ./...
