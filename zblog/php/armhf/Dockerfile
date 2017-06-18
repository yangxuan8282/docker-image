FROM yangxuan8282/rpi-php:7.1.2-alpine

ENV ZBLOG_VERSION=1.5.1
ENV ZBLOG_URL="https://github.com/zblogcn/zblogphp/releases/download/1740/Z-BlogPHP_1_5_1_1740_Zero.zip"

RUN set -x \
  && mkdir -p /usr/src/zblog \
  && apk --update --no-cache add wget bash \
  && wget -O /tmp/zblog.zip "$ZBLOG_URL" \
  && unzip /tmp/zblog.zip -d /usr/src/zblog \ 
  && rm -rf /tmp/*

WORKDIR /var/www/html

COPY entrypoint.sh /usr/local/bin/

EXPOSE 80

ENTRYPOINT ["entrypoint.sh"]
CMD [ "php", "-S", "0000:80", "-t", "/var/www/html" ]
