#!/bin/bash
#
# This script will build the executable and leave it in this directory.
#
# If the first argument to the script is set to "alpine", then it will
# build a binary for Alpine Linux.
#

IMAGE=$([ "$1" == "alpine" ] && echo "yangxuan8282/golang:alpine" || echo "golang")
SRCDIR=/go/src/github.com/ddollar/forego

docker $DOCKER_OPTIONS run -t --rm -v "$PWD":$SRCDIR -e GOOS=${GOOS:-linux} -e GOARCH=${GOARCH:-arm} -w $SRCDIR $IMAGE go build -v .
