BUILD:

```
docker build -t local/freedos .
```

RUN:

```
docker run -d --name freedos -p 6080:6080 local/freedos
```
