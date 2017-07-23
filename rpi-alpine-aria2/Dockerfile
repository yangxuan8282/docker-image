FROM pipill/armhf-alpine
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV ARIA2_VERSION=1.32.0

RUN [ "cross-build-start" ]
RUN apk update \
 && apk upgrade  \
 && apk add bash \
 && apk add curl \
 && apk add tar \
 && rm -rf /var/cache/apk/* \
 && curl -C - -SL https://github.com/q3aql/aria2-static-builds/releases/download/v${ARIA2_VERSION}/aria2-${ARIA2_VERSION}-linux-gnu-arm-rbpi-build1.tar.bz2 -o /tmp/aria2.tar.bz2 \
 && tar xf /tmp/aria2.tar.bz2 -C /tmp \
 && install -Dm755 /tmp/aria2-${ARIA2_VERSION}-linux-gnu-arm-rbpi-build1/aria2c /usr/local/bin \ 
 && rm -rf /tmp \
 && adduser -D aria2 \
 && mkdir -p /etc/aria2 \
 && touch /etc/aria2/session.lock 
RUN [ "cross-build-end" ] 
COPY aria2.conf /etc/aria2/aria2.conf 
VOLUME /home/aria2 /etc/aria2

EXPOSE 6800
CMD set -xe && \
 chown -R aria2:aria2 /home/aria2 && \
 aria2c --conf-path=/etc/aria2/aria2.conf
