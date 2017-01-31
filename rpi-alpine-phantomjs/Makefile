NAME=artifacts/phantomjs-alpine-x86_64.tar.bz2
all: $(NAME)

$(NAME):
	docker build -t phantomjs-alpine . && docker run --rm -i -v `pwd`/artifacts:/artifacts phantomjs-alpine:latest cp /root/phantomjs.tar.bz2 /$(NAME)


clean:
	rm artifacts/* 2>/dev/null || echo "clean"

