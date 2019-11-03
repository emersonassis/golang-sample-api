#!/bin/bash

export GOROOT=/usr/local/go
export PATH_CDM=src/projects.org/sample/sample-api/cmd

go build -v $GOPATH/$PATH_CDM/main.go
./main