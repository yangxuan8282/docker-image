aria2 docker for raspberrypi
---

test on Raspbian Lite 

BT works

base images: hypriot/rpi-alpine-scratch

run with:

```
mkdir -p $HOME/Downloads &&
sudo docker run -d --name=aria2 -p 6800:6800 -v $HOME/Downloads:/home/aria2 yangxuan8282/rpi-alpine-aria2
```
