FROM yangxuan8282/rpi-alpine:3.5
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV GRAFANA_VERSION 4.1.1

RUN apk --no-cache add nodejs

RUN apk update && apk upgrade \
  && apk --no-cache add fontconfig curl libpng-dev freetype-dev \
  && apk --virtual build-deps add build-base go git gcc python musl-dev make nodejs-dev fontconfig-dev \
  && apk update && apk add --no-cache fontconfig curl \
  && mkdir -p /usr/share \
  && cd /usr/share \
  && curl -L https://github.com/yangxuan8282/docker-image/releases/download/2.1.1/phantomjs-2.1.1-alpine-arm.tar.xz | tar xJ \
  && ln -s /usr/share/phantomjs/phantomjs /usr/bin/phantomjs \
  && export GOPATH=/go \
  && mkdir -p $GOPATH/src/github.com/grafana && cd $GOPATH/src/github.com/grafana \
  && git clone https://github.com/grafana/grafana.git -b v${GRAFANA_VERSION} \
  && export GOPATH=/go \
  && cd $GOPATH/src/github.com/grafana/grafana \
  && npm install -g yarn@0.19.0 --registry=http://r.cnpmjs.org \
  && npm install -g grunt-cli@1.2.0 --registry=http://r.cnpmjs.org \
  && go run build.go setup \
  && go run build.go build \
  && yarn install \
  && npm run build \
  && npm uninstall -g yarn \
  && npm uninstall -g grunt-cli \
  && npm cache clear \
  && mv ./bin/grafana-server ./bin/grafana-cli /bin/ \
  && mkdir -p /etc/grafana/json /var/lib/grafana/plugins /var/log/grafana /usr/share/grafana \
  && mv ./public_gen /usr/share/grafana/public \
  && mv ./conf /usr/share/grafana/conf \
  && apk del build-deps && cd / && rm -rf /var/cache/apk/* /usr/local/share/.cache $GOPATH

VOLUME ["/var/lib/grafana", "/var/lib/grafana/plugins", "/var/log/grafana"]

EXPOSE 3000

ENV INFLUXDB_HOST localhost
ENV INFLUXDB_PORT 8086
ENV INFLUXDB_PROTO http
ENV INFLUXDB_USER grafana
ENV INFLUXDB_PASS changeme
ENV GRAFANA_USER admin
ENV GRAFANA_PASS changeme
#ENV GRAFANA_BASE_URL
#ENV FORCE_HOSTNAME

COPY grafana.ini /usr/share/grafana/conf/defaults.ini.tpl
COPY run.sh /run.sh

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["/run.sh"]

HEALTHCHECK --interval=5s --retries=5 --timeout=2s CMD curl -u $GRAFANA_USER:$GRAFANA_PASS 127.0.0.1:3000/api/org 2>/dev/null | grep -q '"id":'
