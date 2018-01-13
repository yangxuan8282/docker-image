FROM base/archlinux

COPY run_novnc /usr/bin/

RUN set -xe \
  && pacman -Syu --noconfirm deepin deepin-extra 

RUN pacman -S --noconfirm tigervnc git net-tools wqy-zenhei firefox \
  && git clone https://github.com/kanaka/noVNC /root/noVNC \
  && git clone https://github.com/kanaka/websockify /root/noVNC/utils/websockify \
  && pacman -R --noconfirm git \
  && rm -rf /var/cache/pacman/pkg/* \
  && chmod +x /usr/bin/run_novnc

ENV USER=root VNCPASSWD=archlinux

RUN echo "$VNCPASSWD\n$VNCPASSWD\nn" | vncpasswd

EXPOSE 6080

CMD ["run_novnc"]
