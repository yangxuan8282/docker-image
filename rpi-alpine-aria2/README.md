aria2 docker for raspberrypi
---

[![](https://images.microbadger.com/badges/image/yangxuan8282/rpi-alpine-aria2.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-aria2 "Get your own image badge on microbadger.com")

test on Raspbian Lite 

### FROM

base images: hypriot/rpi-alpine-scratch

### RUN

```bash
mkdir -p $HOME/Downloads &&
docker run -d --name=aria2 \
  -p 6800:6800 \
  -v $HOME/Downloads:/home/aria2 \
  yangxuan8282/rpi-alpine-aria2
```
