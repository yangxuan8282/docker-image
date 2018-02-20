### TAG

latest, amd64

[![](https://images.microbadger.com/badges/image/yangxuan8282/alpine-xfce4-novnc.svg)](https://microbadger.com/images/yangxuan8282/alpine-xfce4-novnc "Get your own image badge on microbadger.com")

### RUN

```
docker run -d -p 6080:6080 yangxuan8282/alpine-xfce4-novnc:amd64
```

visit YOUR_IP:6080

the default password is `alpinelinux`

to manage docker on host:

```
-v /var/run/docker.sock:/var/run/docker.sock
```

then install docker in container

to run in a local nested X window

run:

```
Xephyr -screen 1024x768 :1 &
docker run -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/alpine-xfce4-novnc:amd64 startxfce4
```

to run multiple container:

```
Xephyr -screen 1024x768 :2 &
docker run -v /tmp:/tmp yangxuan8282/alpine-xfce4-novnc:amd64 startxfce4
```

it is also possible to nested in VNC:

```
sudo Xephyr -screen 800x600 :2 &
sudo docker run -e DISPLAY=:2 -v /tmp:/tmp yangxuan8282/alpine-xfce4-novnc:amd64 startxfce4
```

> need to mount `/tmp/.X11-unix` when start container:

```
docker run -d -p 6080:6080 -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/alpine-xfce4-novnc:amd64
```

to use mmal on raspberry pi:

```
--device /dev/vchiq
```

### TRY

[![Try in PWD](https://github.com/play-with-docker/stacks/raw/cff22438cb4195ace27f9b15784bbb497047afa7/assets/images/button.png)](http://play-with-docker.com?stack=https://raw.githubusercontent.com/yangxuan8282/docker-image/master/alpine-xfce4-novnc/stack.yml)
