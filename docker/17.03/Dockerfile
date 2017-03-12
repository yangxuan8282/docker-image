FROM hypriot/rpi-alpine:3.5

RUN apk add --no-cache \
		ca-certificates \
		curl \
		openssl \
                bash

ENV DOCKER_BUCKET get.docker.com
ENV DOCKER_VERSION 17.03.0-ce
ENV DOCKER_SHA256 3c1c8b46292a199bf818ffad55b1068ecdd7d9d5b70533df5988a970233e82d4

RUN set -x \
	&& curl -fSL "https://${DOCKER_BUCKET}/builds/Linux/armel/docker-${DOCKER_VERSION}.tgz" -o docker.tgz \
	&& echo "${DOCKER_SHA256} *docker.tgz" | sha256sum -c - \
	&& tar -xzvf docker.tgz \
	&& mv docker/* /usr/local/bin/ \
	&& rmdir docker \
	&& rm docker.tgz \
	&& docker -v

COPY docker-entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["sh"]
