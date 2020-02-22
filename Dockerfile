FROM golang:latest

RUN mkdir -p /go/src/practicing-redis-golang

WORKDIR /go/src/practicing-redis-golang

COPY . /go/src/practicing-redis-golang

RUN go get gopkg.in/go-playground/validator.v10
RUN mkdir -p /go/src/practicing-redis-golang/vendor/github.com/go-playground/validator/v10
RUN cp -rf /go/src/gopkg.in/go-playground/validator.v10/* /go/src/practicing-redis-golang/vendor/github.com/go-playground/validator/v10
RUN go install

ENTRYPOINT /go/bin/practicing-redis-golang
