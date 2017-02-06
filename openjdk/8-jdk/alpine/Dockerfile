FROM hypriot/rpi-alpine-scratch:v3.4
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV LANG C.UTF-8

RUN sed -i 's/nl.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update \
  && apk upgrade \
  && apk add bash \
  && rm -rf /var/cache/apk/*

RUN { \
		echo '#!/bin/sh'; \
		echo 'set -e'; \
		echo; \
		echo 'dirname "$(dirname "$(readlink -f "$(which javac || which java)")")"'; \
	} > /usr/local/bin/docker-java-home \
	&& chmod +x /usr/local/bin/docker-java-home
ENV JAVA_HOME /usr/lib/jvm/java-1.8-openjdk
ENV PATH $PATH:/usr/lib/jvm/java-1.8-openjdk/jre/bin:/usr/lib/jvm/java-1.8-openjdk/bin

ENV JAVA_VERSION 8u111
ENV JAVA_ALPINE_VERSION 8.111.14-r1

RUN set -x \
	&& apk add openjdk8="$JAVA_ALPINE_VERSION" \
                --update-cache --repository https://mirrors.ustc.edu.cn/alpine/edge/community/ --allow-untrusted \
	&& [ "$JAVA_HOME" = "$(docker-java-home)" ]
