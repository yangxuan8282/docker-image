resilio sync docker for Raspberry Pi
---

[![](https://images.microbadger.com/badges/image/yangxuan8282/rpi-resilio-sync.svg)](https://microbadger.com/images/yangxuan8282/rpi-resilio-sync "Get your own image badge on microbadger.com")

### RUN

```bash
mkdir -p $HOME/sync/share &&
docker run -d --name Sync \
  -p 8888:8888 -p 55555 \
  -v $HOME/sync/share:/mnt/sync \
  --restart on-failure \
  yangxuan8282/rpi-resilio-sync
```

then visit http://raspberrypi:8888 (or replace `raspberrypi` with your hostname) to setup
