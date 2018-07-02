FROM arm32v6/alpine

RUN apk add --no-cache darkhttpd

EXPOSE 80

VOLUME /www

CMD ["darkhttpd", "/www"]
