FROM yangxuan8282/rpi-alpine-nginx:1.11.10

ENV GOMPLATE_VER v1.1.2
ENV REMARK_VER 0.14.0
ENV WEBROOT /usr/share/nginx/html
ENV TITLE My Title
ENV RATIO 4:3
ENV HIGHLIGHT_LINES false
ENV HIGHLIGHT_STYLE default

ADD https://github.com/hairyhenderson/gomplate/releases/download/$GOMPLATE_VER/gomplate_linux-arm-slim /usr/local/bin/gomplate
RUN chmod a+rx /usr/local/bin/gomplate

ADD https://gnab.github.io/remark/downloads/remark-${REMARK_VER}.min.js /usr/share/nginx/html/
RUN chmod a+r /usr/share/nginx/html/remark-${REMARK_VER}.min.js
COPY styles.css $WEBROOT
COPY index.html.tmpl /
COPY docker-entrypoint.sh /

ENTRYPOINT [ "/docker-entrypoint.sh" ]

CMD ["-g", "daemon off;"]
