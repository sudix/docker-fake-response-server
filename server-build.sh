#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o fake-server go-server/main.go
