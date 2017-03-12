package cmd

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Unknwon/com"
	"github.com/go-xiaohei/pugo/app/compile"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/urfave/cli"
)

var (
	/* serverFlags = append(commonFlags, cli.BoolFlag{
		Name:  "build",
		Usage: "force to build before server starting",
	})*/
	serverFlags = append(commonFlags, cli.StringFlag{
		Name:  "addr",
		Usage: "set http address to start",
		Value: "0.0.0.0:9899",
	})
)

// Server is `server` Command
var Server = cli.Command{
	Name:  "server",
	Usage: "start embedded http server",
	Flags: serverFlags,
	Action: func(cliCtx *cli.Context) error {
		if cliCtx.Bool("debug") {
			printer.EnableLogf = true
		}
		if isSiteAvailable() {
			/*
				ctx, ok := isSiteBuilt()
				if !ok || cliCtx.Bool("build") {
					printer.Trace("build site to server")
					ctx = buildOnce()
				} else {
					printer.Trace("site is ready for server")
				}
				if ctx == nil {
					return nil
				}*/
			watchCtx = buildOnce()
			go runServer(cliCtx.String("addr"))
			watchLoop()
		}
		return nil
	},
}

func isSiteBuilt() (*compile.Context, bool) {
	ctx, err := compile.Payload()
	if err != nil {
		printer.Error("read fail : %v", err)
		return nil, false
	}
	return ctx, com.IsDir(ctx.DstDir())
}

func runServer(addrString string) {
	printer.Print("--- starting server...")
	printer.Trace("server : %v", addrString)

	ctx := watchCtx
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, _ := ctx.Meta.RootURL()
		if u.Path != "" {
			if !strings.HasPrefix(r.URL.Path, u.Path) {
				http.Redirect(w, r, path.Join(u.Path, r.URL.Path), http.StatusTemporaryRedirect)
				return
			}
		}
		dstFiles := url2DstFile(r.URL.Path)
		for _, file := range dstFiles {
			file = path.Join(ctx.Config.OutputDir, file)
			if com.IsFile(file) {
				printer.Logf("serve file %s", file)
				http.ServeFile(w, r, file)
				return
			}
		}
		printer.Logf("serve 404 %s", r.RequestURI)
		http.NotFound(w, r)
	})
	if err := http.ListenAndServe(addrString, nil); err != nil {
		printer.Error("server error : %v", err)
		os.Exit(2)
	}
}

func url2DstFile(urlStr string) []string {
	if path.Ext(urlStr) != "" {
		return []string{urlStr}
	}
	var s []string
	s = append(s, path.Join(urlStr, ".html"))
	s = append(s, path.Join(urlStr, "index.html"))
	return s
}
