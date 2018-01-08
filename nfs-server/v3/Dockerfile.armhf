FROM arm32v6/alpine

COPY entrypoint.sh /entrypoint.sh

RUN apk --no-cache --update add bash nfs-utils \
  && chmod +x /entrypoint.sh

EXPOSE 111 111/udp 2049 2049/udp \
    32765 32765/udp 32766 32766/udp \
    32767 32767/udp 32768 32768/udp

ENTRYPOINT ["/entrypoint.sh"]
