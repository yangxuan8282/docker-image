FROM takashifss/rpi-alpine-base:3.5
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories 

RUN echo "@edge http://mirrors.ustc.edu.cn/alpine/edge/main" >> /etc/apk/repositories && \
    echo "@testing http://mirrors.ustc.edu.cn/alpine/edge/testing" >> /etc/apk/repositories && \
    echo "@community http://mirrors.ustc.edu.cn/alpine/v3.5/community" >> /etc/apk/repositories

RUN apk update && \
    apk upgrade && \
    apk --no-cache add ca-certificates curl wget bash jq gosu@testing && \
    curl -o /usr/bin/envtpl -L https://github.com/yangxuan8282/docker-image/blob/master/envtpl/envtpl?raw=true && \
    chmod a+x /usr/bin/envtpl && \
    rm -rf /var/cache/apk/*

COPY sut /usr/local/sut
