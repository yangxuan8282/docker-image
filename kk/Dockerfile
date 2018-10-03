FROM python:3-slim

RUN set -xe \
  && apt-get update \
  && apt-get install -y gcc \
  && pip3 install --no-cache-dir kk \
  && apt-get purge -y gcc \
  && apt-get autoremove -y \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

CMD ["kk"]
