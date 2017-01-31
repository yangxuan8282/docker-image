FROM hypriot/rpi-alpine-scratch:v3.4
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV SPEED-TEST_VERSION 1.7.1

RUN sed -i 's/nl.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk --no-cache add nodejs

RUN npm install --global speed-test --registry=http://r.cnpmjs.org

CMD ["speed-test"]
