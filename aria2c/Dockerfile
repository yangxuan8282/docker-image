FROM scratch

ENV ARIA2C_VERSION 1.31.0

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY aria2c /aria2c

ENTRYPOINT ["/aria2c"]
