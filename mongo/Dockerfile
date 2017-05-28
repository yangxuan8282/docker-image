FROM pipill/armhf-debian:jessie

# add our user and group first to make sure their IDs get assigned consistently, regardless of whatever dependencies get added
RUN groupadd -r mongodb && useradd -r -g mongodb mongodb

RUN apt-get update \
        && apt-get install -y curl mongodb \
        && rm -rf /var/lib/apt/lists/*

RUN curl -o /usr/local/bin/gosu -SL 'https://github.com/tianon/gosu/releases/download/1.10/gosu-armhf' \
        && chmod +x /usr/local/bin/gosu

ENV MONGO_VERSION 2.4.10

VOLUME /data/db

ADD docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 27017
CMD ["mongod"]

