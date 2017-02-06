first create network

```
docker network create proxy-network
```

edit the `nginx-proxy/docker-compose.yml`

start the `nginx-proxy` with:

```bash
cd nginx-proxy &&
docker-compose up -d
```

then create your `docker-compose.yml` follow the template file `docker-compose.template`,and start them with `docker-compose up -d`

you may need to edit your domain name to a wildcard DNS record on your domain name server(like godaddy)

          from                to
    @.mydomain.com ---> *.mydomain.com




