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

remote X session with ssh tunnel:

run container on server:

```
docker run -d -p 6080:6080 -p 2222:22 -v /var/run/docker.sock:/var/run/docker.sock -v /tmp/.X11-unix:/tmp/.X11-unix yangxuan8282/alpine-xfce4-novnc:amd64
```

run command in VNC:

```
sudo apk --update --no-cache add openssh &&
sudo bash -c 'echo "X11Forwarding yes" >> /etc/ssh/sshd_config' &&
sudo bash -c 'echo "X11UseLocalhost no" >> /etc/ssh/sshd_config' &&
sudo ssh-keygen -A &&
sudo /usr/sbin/sshd -D
```

then on client:

```
Xephyr -screen 800x600 :1 &
DISPLAY=:1.0 ssh -Xf alpine@SERVER_IP -p 2222 xfce4-session
```

> the default ssh passwd is `alpine`

to use mmal on raspberry pi:

```
--device /dev/vchiq
```

### TRY

[![Try in PWD](https://github.com/play-with-docker/stacks/raw/cff22438cb4195ace27f9b15784bbb497047afa7/assets/images/button.png)](http://play-with-docker.com?stack=https://raw.githubusercontent.com/yangxuan8282/docker-image/master/alpine-xfce4-novnc/stack.yml)

![](https://github.com/yangxuan8282/docker-image/raw/master/alpine-xfce4-novnc/chrome_2018-09-23_07-55-15.png)
