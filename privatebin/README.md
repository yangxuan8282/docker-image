[![](https://images.microbadger.com/badges/image/jgeusebroek/privatebin.svg)](https://microbadger.com/images/jgeusebroek/privatebin "Get your own image badge on microbadger.com")
# Docker Privatebin image

A tiny image running [alpine](https://github.com/gliderlabs/docker-alpine) Linux and [Privatebin](https://github.com/PrivateBin/PrivateBin).

## Usage

	docker run --restart=always -d \
		-p 0.0.0.0:80:80 \
		--hostname=privatebin \
		--name=privatebin \
		-v /<host_data_directory>:/privatebin/data \
		-v /<host_cfg_directory>:/privatebin/cfg \
		jgeusebroek/privatebin

On first run it will copy the sample config file if there isn't a config file already.

## Optional environment variables

* `UID` User ID php fpm daemon account (default: 991).
* `GID` Group ID php fpm daemon account (default: 991).

## License

MIT / BSD

## Author Information

[Jeroen Geusebroek](http://jeroengeusebroek.nl/)
