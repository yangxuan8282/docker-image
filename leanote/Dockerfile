FROM pipill/armhf-alpine:edge

ENV LEANOTE_VERSION=2.4

RUN apk add --no-cache --update wget ca-certificates \
    && wget https://jaist.dl.sourceforge.net/project/leanote-bin/${LEANOTE_VERSION}/leanote-linux-arm-v${LEANOTE_VERSION}.bin.tar.gz \
    && tar -zxf leanote-linux-arm-v${LEANOTE_VERSION}.bin.tar.gz -C / \
    && rm -rf /leanote/mongodb_backup \
    && rm leanote-linux-arm-v${LEANOTE_VERSION}.bin.tar.gz \
    && chmod a+x /leanote/bin/run.sh && chmod a+x /leanote/bin/leanote-linux-arm \
    && sed -i '/chmod/d' /leanote/bin/run.sh \
    && apk del --purge wget

VOLUME /leanote/public/upload

EXPOSE 9000

CMD ["/leanote/bin/run.sh"]
