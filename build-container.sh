#!/bin/sh
go get
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o fsw
chmod a+x fsw
docker build -t roffe/kube-fswatch:latest .
docker push roffe/kube-fswatch:latest
