FROM alpine

RUN apk add --no-cache dnsmasq

EXPOSE 53/tcp \
       53/udp \
       67/udp

ENTRYPOINT ["dnsmasq", "--no-daemon", "--user=dnsmasq", "--group=dnsmasq"]
