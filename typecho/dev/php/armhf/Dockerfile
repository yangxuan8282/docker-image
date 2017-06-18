FROM yangxuan8282/rpi-php:7.1.2-alpine

ENV TYPECHO_URL="http://typecho.org/build.tar.gz"

RUN set -x \
  && mkdir -p /usr/src/typecho \
  && apk --update --no-cache add wget bash \
  && wget -O /tmp/typecho.tar.gz "$TYPECHO_URL" \
  && tar -xzf /tmp/typecho.tar.gz -C /usr/src/typecho/ --strip-components=1 \ 
  && rm -rf /tmp/*

WORKDIR /var/www/html

COPY entrypoint.sh /usr/local/bin/

EXPOSE 80

ENTRYPOINT ["entrypoint.sh"]
CMD [ "php", "-S", "0000:80", "-t", "/var/www/html" ]
