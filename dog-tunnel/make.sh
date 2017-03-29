#!/bin/bash
#
# This script will build the executable and leave it in this directory.
#
# If the first argument to the script is set to "alpine", then it will
# build a binary for Alpine Linux.
#

IMAGE=golang:alpine
SRCDIR=/go/src/github.com/vzex/dog-tunnel
DTUNNEL_VERSION=lite_v1.30

git clone https://github.com/vzex/dog-tunnel -b ${DTUNNEL_VERSION} ./src
docker $DOCKER_OPTIONS run -t --rm -v "$PWD"/src:$SRCDIR -v "$PWD"/build.sh:/build.sh -e DTUNNEL_VERSION=${DTUNNEL_VERSION} -e SRCDIR=${SRCDIR} -e GOOS=${GOOS:-linux} -e GOARCH=${GOARCH:-amd64} -e CGO_ENABLED=${CGO_ENABLED:-0} -w $SRCDIR $IMAGE sh /build.sh
#docker $DOCKER_OPTIONS run -t --rm -v "$PWD"/src:$SRCDIR -v "$PWD"/build.sh:/build.sh -e DTUNNEL_VERSION=${DTUNNEL_VERSION} -e SRCDIR=${SRCDIR} -e GOOS=${GOOS:-linux} -e GOARCH=${GOARCH:-arm} -e CGO_ENABLED=${CGO_ENABLED:-0} -w $SRCDIR $IMAGE sh /build.sh
