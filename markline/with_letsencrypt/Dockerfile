FROM hypriot/rpi-alpine-scratch:v3.4
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV MARKLINE_VERSION 0.6.0

RUN sed -i 's/nl.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk --no-cache add nodejs

RUN npm install -g markline --registry=http://r.cnpmjs.org

EXPOSE 8000

VOLUME ["/root"]

WORKDIR /root

CMD ["markline", "server", "life.md"]
