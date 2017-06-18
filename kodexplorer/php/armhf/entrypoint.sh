#!/bin/bash

set -e

if [ "$1" = 'php' ] && [ "$(id -u)" = '0' ]; then
    chown -R www-data /var/www/html
    chmod -R 777 /var/www/html/
fi

if [ ! -e '/var/www/html/index.php' ]; then
    cp -a /usr/src/kodexplorer/* /var/www/html/
fi

exec "$@"
