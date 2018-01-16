FROM debian
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV DEBIAN_FRONTEND=noninteractive \
    DISPLAY=:1 \
    VNC_PORT=5901 \
    VNC_RESOLUTION=1024x768 \
    VNC_COL_DEPTH=24 

COPY run_vnc /usr/bin/
COPY keyboard /etc/default/keyboard

RUN set -xe \
  && apt-get update && apt-get install -y dirmngr \
  && echo "deb http://archive.raspberrypi.org/debian/ stretch main ui" | tee /etc/apt/sources.list.d/raspi.list \
  && apt-key adv --keyserver keyserver.ubuntu.com --recv 82B129927FA3303E \
  && apt-get update && apt-get upgrade -y \
  && apt-get install -y raspberrypi-ui-mods lxterminal rpi-chromium-mods tightvncserver htop \
  && useradd -g sudo -ms /bin/bash pi \
  && echo 'pi:raspberry' | chpasswd -e \
  && echo 'pi ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers \
  && sed -i 's/#force_color_prompt=yes/force_color_prompt=yes/g' /home/pi/.bashrc \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && chmod +x /usr/bin/run_vnc 

ENV USER=pi \
    VNC_PASSWD=raspberry \
    PASSWD_PATH="$HOME/.vnc/passwd" 

EXPOSE $VNC_PORT

CMD ["run_vnc"]
