FROM alpine

ENV GOTTY_VERSION=1.0.1 \
    TERM=xterm

RUN apk --update --no-cache add bash wget \
  && wget https://github.com/yudai/gotty/releases/download/v"$GOTTY_VERSION"/gotty_linux_386.tar.gz -O - | tar -xz -C /usr/bin/ \
  && apk del wget

EXPOSE 8080

ENTRYPOINT ["gotty"]

CMD ["-w", "bash"]
