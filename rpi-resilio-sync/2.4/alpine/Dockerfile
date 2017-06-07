FROM pipill/armhf-alpine:3.4
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV RESILIO_VERSION=2.4.4

ENV RESILIO_URL=https://download-cdn.resilio.com/${RESILIO_VERSION}/linux-armhf/resilio-sync_armhf.tar.gz

RUN apk update \
 && apk upgrade \
 && apk add curl \
 && apk add bash \
 && apk add libc6-compat \
 && rm -rf /var/cache/apk/* \
 && curl -C - -SL ${RESILIO_URL} -o /tmp/sync.tgz \
 && tar xf /tmp/sync.tgz -C /usr/bin rslsync \ 
 && rm -f /tmp/sync.tgz

COPY sync.conf.default /etc/
COPY run_sync /usr/bin/

EXPOSE 8888 
EXPOSE 55555

VOLUME /mnt/sync

ENTRYPOINT ["run_sync"]
CMD ["--config", "/mnt/sync/sync.conf"]
