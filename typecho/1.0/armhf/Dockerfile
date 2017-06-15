FROM yangxuan8282/rpi-php:7.1-apache

ENV TYPECHO_VERSION=1.0
ENV TYPECHO_URL="https://github.com/typecho/typecho/releases/download/v1.0-14.10.10-release/1.0.14.10.10.-release.tar.gz"

RUN set -x \
  && mkdir -p /usr/src/typecho \
  && apt-get update && apt-get install -y --no-install-recommends ca-certificates wget && rm -rf /var/lib/apt/lists/* \
  && wget -O /tmp/typecho.tar.gz "$TYPECHO_URL" \
  && tar -xzf /tmp/typecho.tar.gz -C /usr/src/typecho/ --strip-components=1 \
  && apt-get purge -y --auto-remove ca-certificates wget \
  && rm -rf /var/cache/apk/* \
  && rm -rf /tmp/*

WORKDIR /var/www/html

COPY entrypoint.sh /usr/local/bin/

EXPOSE 80

ENTRYPOINT ["entrypoint.sh"]
CMD ["apache2-foreground"]
