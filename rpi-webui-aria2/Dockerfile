FROM hypriot/rpi-alpine-scratch
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

RUN apk update \
 && apk upgrade \
 && apk add curl \
 && apk add nginx \
 && mkdir -p /run/nginx \
 && rm -rf /var/cache/apk/* \
 && mkdir -p /tmp \
 && curl -C - -SL https://github.com/ziahamza/webui-aria2/archive/master.zip -o /tmp/webui-aria2-master.zip \
 && unzip /tmp/webui-aria2-master.zip -d /tmp \
 && cp -a /tmp/webui-aria2-master/* /usr/share/nginx/html \
 && rm -rf /tmp

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
