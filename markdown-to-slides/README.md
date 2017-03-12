# hairyhenderson/remarkjs

nginx-based hosting for _external_ [remarkjs] slides.

## Usage

- mount your remark-formatted slides file to `/slides.md`

```console
$ docker run -d -p 8080:80 -v $(pwd)/slides.md:/slides.md hairyhenderson/remarkjs
$ open http://localhost:8080
```

## Development Usage

Enable development mode by setting the `DEV_MODE` environment variable to any value. Then, mount your slides at `/usr/share/nginx/html/slides.md`.

```console
$ docker run -p 8080:80 -v $(pwd)/slides.md:/usr/share/nginx/html/slides.md -e DEV_MODE=true hairyhenderson/remarkjs
$ open http://localhost:8080
```

## Overriding Styles

To customize the look and feel, you can provide your own stylesheet. If placed at `/styles.css`, it wi
ll be copied into `$WEBROOT/styles.css` upon startup.  While in development, mounting directly to `/us
r/share/nginx/html/styles.css` will ensure local changes to the CSS will be reflected upon next refres
h.

[remarkjs]: http://remarkjs.com/#1
