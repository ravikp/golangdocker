#!/bin/bash

docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:latest sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'


go get github.com/tools/godep
