# Resilio Sync
#
# VERSION               0.1
#

FROM frolvlad/alpine-glibc
MAINTAINER Resilio Inc. <support@resilio.com>
LABEL com.resilio.version="2.5.12"

ADD https://download-cdn.resilio.com/2.5.12/linux-x64/resilio-sync_x64.tar.gz /tmp/sync.tgz
RUN tar -xf /tmp/sync.tgz -C /usr/bin rslsync && rm -f /tmp/sync.tgz
RUN apk --no-cache --update add bash

COPY sync.conf.default /etc/
COPY run_sync /usr/bin/

EXPOSE 8888
EXPOSE 55555

VOLUME /mnt/sync

ENTRYPOINT ["run_sync"]
CMD ["--config", "/mnt/sync/sync.conf"]

