FROM debian
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

COPY run_vnc /usr/bin/
COPY keyboard /etc/default/keyboard

RUN set -xe \
  && apt-get update && apt-get install -y dirmngr \
  && echo "deb http://archive.raspberrypi.org/debian/ stretch main ui" | tee /etc/apt/sources.list.d/raspi.list \
  && apt-key adv --keyserver keyserver.ubuntu.com --recv 82B129927FA3303E \
  && apt-get update && apt-get upgrade -y \
  && apt-get install -y raspberrypi-ui-mods lxterminal rpi-chromium-mods tightvncserver \
  && rm -rf /var/lib/apt/lists/* \
  && chmod +x run_vnc

ENV USER=root VNCPASSWD=raspberry

RUN echo "$VNCPASSWD\n$VNCPASSWD\nn" | vncpasswd

EXPOSE 5901

CMD ["run_vnc"]
