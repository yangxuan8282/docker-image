#!/bin/sh

addgroup -g ${GID} privatebin && adduser -h /privatebin -s /bin/sh -D -G privatebin -u ${UID} privatebin
touch /var/run/php-fpm.sock

if [ ! -f /privatebin/cfg/conf.ini ]; then
	cp /privatebin/conf.ini.sample /privatebin/cfg/conf.ini
fi

chown -R privatebin:privatebin /privatebin /var/run/php-fpm.sock /var/lib/nginx /tmp
supervisord -c /usr/local/etc/supervisord.conf
