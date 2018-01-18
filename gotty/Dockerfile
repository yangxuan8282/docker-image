FROM alpine

ADD gotty /usr/bin/

RUN apk --update --no-cache add bash

ENV TERM=xterm

EXPOSE 8080

ENTRYPOINT ["gotty"]

CMD ["-w", "bash"]
