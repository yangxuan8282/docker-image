FROM yangxuan8282/rpi-alpine-docker:17.03-dind

ADD drone-docker /bin/
ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "/bin/drone-docker"]
