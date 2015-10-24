#!/bin/bash

set -e

export GOPATH=$(pwd):$GOPATH

go get github.com/garyburd/redigo/redis
go get github.com/go-martini/martini
go get gopkg.in/alecthomas/kingpin.v1
go get github.com/oguzbilgic/pandik

if [ ! -z "$SKIPTESTS" ]; then
  go test flapjack
fi

go build -x -o libexec/httpbroker libexec/httpbroker.go

if [ ! -z "$CROSSCOMPILE" ]; then
  GOOS=linux GOARCH=amd64 CGOENABLED=0 go build -x -o libexec/httpbroker.linux_amd64 libexec/httpbroker.go
  GOOS=linux GOARCH=386 CGOENABLED=0 go build -x -o libexec/httpbroker.linux_386 libexec/httpbroker.go
fi
