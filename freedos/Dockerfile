FROM alpine:edge

ADD http://www.freedos.org/download/download/FD12CD.iso /
COPY run_qemu /

RUN set -xe \
  && echo "http://dl.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
  && apk --update --no-cache add qemu-img qemu-system-i386 novnc bash \
  && ln -s /usr/share/novnc/vnc_lite.html /usr/share/novnc/index.html

EXPOSE 6080

CMD ["/run_qemu"]
