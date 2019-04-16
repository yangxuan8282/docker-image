FROM alpine:edge

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
  && apk --update --no-cache add tar xz libguestfs \
  && apk --no-cache add ca-certificates wget && update-ca-certificates \
  && wget http://download.libguestfs.org/binaries/appliance/appliance-1.40.1.tar.xz \
  && mkdir -p /usr/lib/guestfs \
  && tar xf *.tar.xz -C /usr/lib/guestfs \
  && rm -f *.tar.xz
