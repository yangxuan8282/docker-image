# run "./build.sh alpine" first to generate envtpl
FROM alpine:3.5
COPY envtpl .
RUN mv envtpl /usr/local/bin
ENTRYPOINT [ "envtpl" ]

