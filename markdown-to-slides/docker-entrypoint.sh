#!/bin/ash
set -e

if [ "${1:0:1}" = '-' ]; then
  set -- nginx "$@"
fi

if [ "$1" = 'nginx' ]; then
  gomplate < /index.html.tmpl > $WEBROOT/index.html
  if [ -z "$DEV_MODE" ]; then
    if [ -f /slides.md.tmpl ]; then
      gomplate < /slides.md.tmpl > $WEBROOT/slides.md
    else
      cp /slides.md $WEBROOT/slides.md
    fi

    if [ -f /styles.css ]; then
      cp /styles.css $WEBROOT/styles.css
    fi

  else
    echo "-- Starting with dev mode enabled"
  fi
fi

exec "$@"
