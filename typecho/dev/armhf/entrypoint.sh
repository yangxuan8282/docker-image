#!/bin/bash

set -e

if [ "$1" = 'apache2-foreground' ] && [ "$(id -u)" = '0' ]; then
    chown -R www-data /var/www/html
    chmod -R 777 /var/www/html/
fi

if [ ! -e '/var/www/html/index.php' ]; then
    su - www-data -s /bin/bash -c 'cp -a /usr/src/typecho/* /var/www/html/'
fi

exec "$@"
