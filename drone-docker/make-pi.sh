#!/bin/bash
#
# This script will build the executable and leave it in this directory.
#
#

IMAGE=yangxuan8282/golang:alpine
SRCDIR=/go/src/github.com/drone-plugins/drone-docker

docker $DOCKER_OPTIONS run -t --rm -v "$PWD":$SRCDIR -e GOOS=${GOOS:-linux} -e GOARCH=${GOARCH:-arm} -e CGO_ENABLED=${CGO_ENABLED:-0} -w $SRCDIR $IMAGE go build .
