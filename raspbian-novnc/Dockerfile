FROM debian
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

COPY run_novnc /usr/bin/
COPY keyboard /etc/default/keyboard

ENV DEBIAN_FRONTEND=noninteractive

RUN set -xe \
  && apt-get update && apt-get install -y dirmngr \
  && echo "deb http://archive.raspberrypi.org/debian/ stretch main ui" | tee /etc/apt/sources.list.d/raspi.list \
  && apt-key adv --keyserver keyserver.ubuntu.com --recv 82B129927FA3303E \
  && apt-get update && apt-get upgrade -y \
  && apt-get install -y raspberrypi-ui-mods lxterminal firefox-esr tightvncserver net-tools wget htop \
  && useradd -g sudo -ms /bin/bash pi \
  && echo "pi:raspberry" | chpasswd \
  && echo "pi ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
  && sed -i 's/#force_color_prompt=yes/force_color_prompt=yes/g' /home/pi/.bashrc \
  && rm -rf /var/lib/apt/lists/* \
  && chmod +x /usr/bin/run_novnc

USER pi

ENV USER=pi \
    DISPLAY=:1 \
    HOME=/home/pi \
    TERM=xterm \
    SHELL=/bin/bash \
    VNC_PASSWD=raspberry \
    VNC_PORT=5901 \
    VNC_RESOLUTION=1024x768 \
    VNC_COL_DEPTH=24 \
    NOVNC_PORT=6080 \
    NOVNC_HOME=/home/pi/noVNC

RUN set -xe \
  && mkdir -p $NOVNC_HOME/utils/websockify \
  && wget -qO- https://github.com/novnc/noVNC/archive/v0.6.2.tar.gz | tar xz --strip 1 -C $NOVNC_HOME \
  && wget -qO- https://github.com/novnc/websockify/archive/v0.6.1.tar.gz | tar xz --strip 1 -C $NOVNC_HOME/utils/websockify \
  && chmod +x -v $NOVNC_HOME/utils/*.sh \
  && ln -s $NOVNC_HOME/vnc_auto.html $NOVNC_HOME/index.html \
  && sudo apt-get purge -y wget \
  && sudo apt-get autoremove -y

WORKDIR $HOME
EXPOSE $VNC_PORT $NOVNC_PORT

CMD ["run_novnc"]
