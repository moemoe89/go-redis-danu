FROM golang:latest

RUN mkdir -p /go/src/practicing-redis-golang

WORKDIR /go/src/practicing-redis-golang

COPY . /go/src/practicing-redis-golang

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-redis-golang
