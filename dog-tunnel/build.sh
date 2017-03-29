#!/bin/sh

apk add --no-cache bash build-base git
go get github.com/cznic/zappy
go get github.com/klauspost/reedsolomon
cd ${SRCDIR}
go build .
