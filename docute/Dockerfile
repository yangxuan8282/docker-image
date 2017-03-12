FROM yangxuan8282/rpi-alpine-node:6-alpine
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

#RUN apk --no-cache add nodejs

RUN npm i -g docute-cli --registry=http://r.cnpmjs.org \
  && mkdir -p /docs

COPY docs/* /docs/

EXPOSE 8080

VOLUME ["/docs/README.md"]

#RUN docute init /docs

CMD ["docute"]
