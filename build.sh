#!/bin/sh

VETSION=latest

GOOS=linux GOARCH=amd64 go build

docker build -t gin-api:$VETSION .
#docker push gin-api:$VETSION