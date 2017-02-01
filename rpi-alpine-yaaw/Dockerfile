FROM yangxuan8282/rpi-alpine-aria2
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

RUN apk update \
 && apk upgrade \
 && apk add curl \
 && apk add nginx \
 && mkdir -p /run/nginx \
 && rm -rf /var/cache/apk/* \
 && mkdir -p /tmp \
 && curl -C - -SL https://github.com/binux/yaaw/archive/master.zip -o /tmp/yaaw-master.zip \
 && unzip /tmp/yaaw-master.zip -d /tmp \
 && cp -a /tmp/yaaw-master/* /usr/share/nginx/html \ 
 && rm -rf /tmp 

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
