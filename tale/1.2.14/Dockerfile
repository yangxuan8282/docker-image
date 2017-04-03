FROM openjdk:8-jdk

ENV TALE_VERSION=1.2.14

WORKDIR /usr/src/myapp

RUN apt-get update && apt-get install -y wget zip sqlite3 \
  && wget http://7xls9k.dl1.z0.glb.clouddn.com/tale.zip && unzip -o tale.zip \
  && rm -f tale.zip \
  && rm -rf /var/cache/apk/*

ENV TALE_HOME /var/tale_home
ENV TALE_SLAVE_AGENT_PORT 9000

VOLUME /var/tale_home

EXPOSE 9000

CMD ["java",  "-jar",  "/usr/src/myapp/tale/tale-1.2.14.jar"]
