FROM php:7.1-alpine

ENV H5AI_VERSION=0.29.0
ENV H5AI_URL="https://release.larsjung.de/h5ai/h5ai-0.29.0.zip"

RUN set -x \
  && mkdir -p /usr/src/h5ai \
  && apk --update --no-cache add wget bash \
  && wget -O /tmp/h5ai.zip "$H5AI_URL" \
  && unzip /tmp/h5ai.zip -d /usr/src/h5ai \
  && rm -rf /tmp/*

COPY router.php /usr/src/h5ai/_h5ai/

WORKDIR /var/www/html

COPY entrypoint.sh /usr/local/bin/

EXPOSE 80

ENTRYPOINT ["entrypoint.sh"]
CMD [ "php", "-S", "0000:80", "_h5ai/router.php" ]
