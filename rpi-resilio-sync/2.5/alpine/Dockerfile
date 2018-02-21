FROM yangxuan8282/armhf-alpine-glibc

ENV RESILIO_VERSION=2.5.12
ENV RESILIO_URL=https://download-cdn.resilio.com/${RESILIO_VERSION}/linux-armhf/resilio-sync_armhf.tar.gz

RUN set -xe \
  && apk --update --no-cache add bash curl \
  && curl ${RESILIO_URL} | tar -xz -C /usr/bin rslsync \
  && apk del curl

COPY sync.conf.default /etc/
COPY run_sync /usr/bin/

EXPOSE 8888
EXPOSE 55555

VOLUME /mnt/sync

ENTRYPOINT ["run_sync"]
CMD ["--config", "/mnt/sync/sync.conf"]
