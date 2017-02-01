grafana docker for Raspberry Pi
---

[![](https://images.microbadger.com/badges/image/yangxuan8282/rpi-alpine-grafana.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-grafana "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/yangxuan8282/rpi-alpine-grafana.svg)](https://microbadger.com/images/yangxuan8282/rpi-alpine-grafana "Get your own version badge on microbadger.com")

### FROM

steal from [appcelerator/docker-grafana](https://github.com/appcelerator/docker-grafana)

base image: yangxuan8282/rpi-alpine:3.5

thanks for help from [Nicolas Degory](https://github.com/ndegory)

### RUN

```bash
docker run -d --name=grafana -p 3000:3000 appcelerator/grafana
```

then visit http://raspberrypi:3000 (or replace raspberrypi with your hostname)

