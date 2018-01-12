FROM arm32v7/debian
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

COPY run_vnc /usr/bin/

RUN set -xe \
  && apt-get update && apt-get install -y dirmngr \
  && rm -f /etc/apt/sources.list \
  && echo "deb http://mirrors.ustc.edu.cn/raspbian/raspbian/ stretch main contrib non-free rpi" | tee /etc/apt/sources.list \
  && echo "deb http://mirrors.ustc.edu.cn/archive.raspberrypi.org/debian/ stretch main ui" | tee /etc/apt/sources.list.d/raspi.list \
  && apt-key adv --keyserver keyserver.ubuntu.com --recv 82B129927FA3303E \
  && apt-key adv --keyserver keyserver.ubuntu.com --recv 9165938D90FDDD2E \
  && apt-get update && apt-get upgrade -y \
  && apt-get install -y raspberrypi-ui-mods lxterminal rpi-chromium-mods tightvncserver \
  && rm -rf /var/lib/apt/lists/* \
  && chmod +x /usr/bin/run_vnc

ENV USER=root VNCPASSWD=raspberry

RUN echo "$VNCPASSWD\n$VNCPASSWD\nn" | vncpasswd

EXPOSE 5901

CMD ["run_vnc"]
