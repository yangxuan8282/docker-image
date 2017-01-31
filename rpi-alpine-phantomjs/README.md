phantomjs docker for Raspberry Pi
---

[![](https://images.microbadger.com/badges/image/yangxuan8282/rpi-alpine-phantomjs.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-phantomjs "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/yangxuan8282/rpi-alpine-phantomjs.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-phantomjs "Get your own version badge on microbadger.com")

### FROM

steal from [Overbryd/docker-phantomjs-alpine](https://github.com/Overbryd/docker-phantomjs-alpine)

base images: hypriot/rpi-alpine-scratch:v3.3

### BUILD

-   2 GB swap

- 4-5 hours

- 1.7 GB image

### RUN

extract binary file:

```bash
docker run -t yangxuan8282/rpi-alpine-phantomjs cat /root/phantomjs.tar.bz2 > phantomjs.tar.bz2
```

Include the binary in your Alpine Dockerfile like this:

```bash
RUN apk update && apk add --no-cache fontconfig && \
  mkdir -p /usr/share && \
  cd /usr/share \
  && curl -L URL | tar xj \
  && ln -s /usr/share/phantomjs/phantomjs /usr/bin/phantomjs \
  && phantomjs --version
```
