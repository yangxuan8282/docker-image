### TAG

armhf, latest

[![](https://images.microbadger.com/badges/image/yangxuan8282/pixel-novnc.svg)](https://microbadger.com/images/yangxuan8282/pixel-novnc "Get your own image badge on microbadger.com")

amd64

[![](https://images.microbadger.com/badges/image/yangxuan8282/pixel-novnc:amd64.svg)](https://microbadger.com/images/yangxuan8282/pixel-novnc:amd64 "Get your own image badge on microbadger.com")

### RUN

for armhf (installed rpi-chromium-mods) :

```
docker run -d -p 2018:6080 yangxuan8282/pixel-novnc
```

for amd64 (installed firefox-esr) :

```
docker run -d -p 2018:6080 yangxuan8282/pixel-novnc:amd64
```

the file manager PCManFM seems not really stable in amd64 version container, it will crash when double click the icons in right panel, so you have to operate from left panel, but armhf version container doesn't have this issues 

then visit http://YOUR_IP:2018 , the default vnc password is `raspberry`

to manage docker on host:

```
-v /var/run/docker.sock:/var/run/docker.sock
```

then install docker in container

to run in a local nested X window

run:

```
Xephyr -screen 1024x768 :1 &
docker run -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/pixel-novnc:amd64 startlxde-pi
```

to simulate multiple monitors:

```
Xephyr -screen 640x480 -screen 640x480 +xinerama :1 &
docker run -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/pixel-novnc:amd64 startlxde-pi
```

to run multiple container:

```
Xephyr -screen 1024x768 :2 &
docker run -v /tmp:/tmp yangxuan8282/pixel-novnc:amd64 startlxde-pi
```

it is also possible to nested in VNC:

```
sudo Xephyr -screen 800x600 :2 &
sudo docker run -e DISPLAY=:2 -v /tmp:/tmp yangxuan8282/pixel-novnc:amd64 startlxde-pi
```

> need to mount `/tmp/.X11-unix` when start container:

```
docker run -d -p 6080:6080 -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/pixel-novnc:amd64
```

to use mmal on raspberry pi:

```
--device /dev/vchiq
```

try it now:

[![Try in PWD](https://github.com/play-with-docker/stacks/raw/cff22438cb4195ace27f9b15784bbb497047afa7/assets/images/button.png)](http://play-with-docker.com?stack=https://raw.githubusercontent.com/yangxuan8282/docker-image/master/raspbian-novnc/stack.yml)

to publish VNC port:

```
-p 5901:5901
```

to set VNC password:

```
-e VNC_PASSWD=foooobar
```

to change resolution:

```
-e VNC_RESOLUTION=1024x768
```

to increase the size of /dev/shm:

```
--shm-size=256m
```

### Dockerfile

[Dockerfile.armhf](https://github.com/yangxuan8282/docker-image/blob/master/raspbian-novnc/Dockerfile.armhf)

[Dockerfile](https://github.com/yangxuan8282/docker-image/blob/master/raspbian-novnc/Dockerfile)

### FROM

arm32v7/debian

debian
