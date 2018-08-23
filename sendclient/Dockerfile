FROM alpine

RUN set -xe \
  && apk --update --no-cache add python3 gcc g++ \
  && pip3 install sendclient

CMD ["send-cli"]
