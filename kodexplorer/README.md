### Tags

3.4.6, 3.4.6-amd64, amd64, latest

[![](https://images.microbadger.com/badges/image/yangxuan8282/kodexplorer.svg)](https://microbadger.com/images/yangxuan8282/kodexplorer "Get your own image badge on microbadger.com")

3.4.6-arm, arm 

[![](https://images.microbadger.com/badges/image/yangxuan8282/kodexplorer:arm.svg)](https://microbadger.com/images/yangxuan8282/kodexplorer:arm "Get your own image badge on microbadger.com")

### RUN

```
docker run -d -p 80:80 --name kodexplorer -v "$PWD":/var/www/html yangxuan8282/kodexplorer
```

### TRY

[![Try in PWD](https://github.com/play-with-docker/stacks/raw/cff22438cb4195ace27f9b15784bbb497047afa7/assets/images/button.png)](http://play-with-docker.com/?stack=https://gist.github.com/yangxuan8282/2f64a4edb2fd6f692b1c8e437ff68468/raw/508435a4fd35e5f378ae7598cfcd67baff24ef62/stack.yml)

### Dockerfile

https://github.com/yangxuan8282/docker-image/tree/master/kodexplorer

### Docker Compose

```
version: "2"

services:
  kodexplorer:
    image: yangxuan8282/kodexplorer
    restart: "always"
    ports:
      - "80:80"
    volumes:
      - "./html:/var/www/html"
```
