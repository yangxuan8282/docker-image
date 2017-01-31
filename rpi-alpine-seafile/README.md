seafile docker for Raspberry Pi
---

[![](https://images.microbadger.com/badges/image/yangxuan8282/rpi-alpine-seafile.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-seafile "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/yangxuan8282/rpi-alpine-seafile.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-seafile "Get your own version badge on microbadger.com")

### FROM

steal from [VGoshev/seafile-docker](https://github.com/VGoshev/seafile-docker)

base images: hypriot/rpi-alpine-scratch:v3.4

### RUN

```bash
install -dm777 $HOME/seafile &&
docker run --name seafile \
  -v $HOME/seafile:/home/seafile \
  -p 192.168.8.103:8000:8000 \
  -p 192.168.8.103:8082:8082 \
  -ti yangxuan8282/rpi-alpine-seafile
```

>replace `192.168.8.103` with your raspberry pi IP address 

then configure admin email and password

web: visit http://raspberrypi:8000 (or replace `raspberrypi` with your hostname/IP address)

app: http://raspberrypi:8000, email, password (or replace `raspberrypi` with your hostname/IP address)
