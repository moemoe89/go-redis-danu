FROM golang:latest

RUN mkdir -p /go/src/go-redis-danu

WORKDIR /go/src/go-redis-danu

COPY . /go/src/go-redis-danu

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-redis-danu
