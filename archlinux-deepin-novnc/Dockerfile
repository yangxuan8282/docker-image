FROM base/archlinux

COPY run_novnc /usr/bin/
ENV VNC_RESOLUTION="1920x1080"
RUN set -xe \
  && pacman -Syu --noconfirm deepin deepin-extra 

RUN pacman -S --noconfirm tigervnc git net-tools wqy-zenhei firefox \
  && git clone https://github.com/kanaka/noVNC /root/noVNC \
  && git clone https://github.com/kanaka/websockify /root/noVNC/utils/websockify \
  && pacman -R --noconfirm git \
  && rm -rf /var/cache/pacman/pkg/* \
  && chmod +x /usr/bin/run_novnc

EXPOSE 6080

CMD ["run_novnc"]
