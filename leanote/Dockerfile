FROM hypriot/rpi-alpine-scratch:v3.3

RUN sed -i 's/nl.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

ENV LEANOTE_VERSION 2.3

RUN apk --update add curl && \
    curl -L http://sourceforge.net/projects/leanote-bin/files/${LEANOTE_VERSION}/leanote-linux-arm-v${LEANOTE_VERSION}.bin.tar.gz/download >> \
    /usr/local/leanote-linux-arm.bin.tar.gz && \
    apk del --purge curl && \
    rm -rf /var/cache/apk/*

RUN tar -xzvf /usr/local/leanote-linux-arm.bin.tar.gz -C /usr/local
RUN chmod +x /usr/local/leanote/bin/run.sh
RUN hash=$(< /dev/urandom tr -dc A-Za-z0-9 | head -c${1:-64};echo;); \
    sed -i "s/app.secret=.*$/app.secret=$hash #/" /usr/local/leanote/conf/app.conf; \
    sed -i "s/db.host=.*$/db.host=db/" /usr/local/leanote/conf/app.conf; \
    sed -i "s/site.url=.*$/site.url=\${SITE_URL} /" /usr/local/leanote/conf/app.conf;

EXPOSE 9000
WORKDIR  /usr/local/leanote/bin
ENTRYPOINT ["sh", "run.sh"]
