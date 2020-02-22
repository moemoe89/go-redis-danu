[![Build Status](https://travis-ci.org/moemoe89/practicing-redis-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-redis-golang)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/practicing-redis-golang/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/practicing-redis-golang?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-redis-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-redis-golang)

# PRACTICING-REDIS-GOLANG #

Practicing Redis Using Golang (Gin Gonic Framework) as Programming Language, Redis as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-redis-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup Redis <https://www.redis.io/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-redis-golang
```

## Running Application
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ go build
```
Run
```
$ go run main.go
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```