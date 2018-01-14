FROM ubuntu:17.10

ENV DEBIAN_FRONTEND noninteractive

COPY run_novnc /usr/bin/
COPY keyboard /etc/default/keyboard

RUN set -xe \
  && add-apt-repository ppa:leaeasy/dde \
  && apt-get update && apt-get upgrade -y \
  && apt-get install -y dde 

RUN apt-get install -y tightvnc git net-tools fonts-wqy-zenhei firefox \
  && git clone https://github.com/kanaka/noVNC /root/noVNC \
  && git clone https://github.com/kanaka/websockify /root/noVNC/utils/websockify \
  && apt-get purge -y git \
  && rm -rf /var/lib/apt/lists/* \
  && chmod +x /usr/bin/run_novnc

ENV USER=root VNCPASSWD=raspberry

RUN echo "$VNCPASSWD\n$VNCPASSWD\nn" | vncpasswd

EXPOSE 6080

CMD ["run_novnc"]
