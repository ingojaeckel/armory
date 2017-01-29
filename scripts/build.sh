#!/bin/sh
go version
go test -v -cover -run Unit
CGO_ENABLED=0 GOOS=linux go build -a -o app
file app
du -hs app
