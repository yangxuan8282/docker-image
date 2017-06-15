FROM php:7.1-apache

ENV ZBLOG_VERSION=1.5.1
ENV ZBLOG_URL="https://github.com/zblogcn/zblogphp/releases/download/1740/Z-BlogPHP_1_5_1_1740_Zero.zip"

RUN set -x \
  && mkdir -p /usr/src/zblog \
  && apt-get update && apt-get install -y --no-install-recommends ca-certificates wget unzip && rm -rf /var/lib/apt/lists/* \
  && wget -O /tmp/zblog.zip "$ZBLOG_URL" \
  && unzip /tmp/zblog.zip -d /usr/src/zblog \
  && apt-get purge -y --auto-remove ca-certificates wget unzip\
  && rm -rf /var/cache/apk/* \
  && rm -rf /tmp/*

WORKDIR /var/www/html

COPY entrypoint.sh /usr/local/bin/

EXPOSE 80

ENTRYPOINT ["entrypoint.sh"]
CMD ["apache2-foreground"]  
