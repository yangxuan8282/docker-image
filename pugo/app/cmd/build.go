package cmd

import (
	"sync"
	"time"

	"github.com/go-xiaohei/pugo/app/compile"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/helper/watcher"
	"github.com/urfave/cli"
)

var (
	buildFlags = append(commonFlags, cli.BoolFlag{
		Name:  "watch",
		Usage: "watch source to re-compile",
	})
	watchCtx *compile.Context
	// watchingExt sets the suffix that watching to
	watchingExt = []string{".md", ".toml", ".html", ".css", ".js", ".jpg", ".png", ".gif"}
	// watchScheduleTime sets watching timer duration
	watchScheduleTime int64
	watchLock         sync.Mutex
)

// Build is 'build' command
var Build = cli.Command{
	Name:  "build",
	Usage: "build your contents to static site",
	Flags: buildFlags,
	Action: func(cliCtx *cli.Context) error {
		if cliCtx.Bool("debug") {
			printer.EnableLogf = true
		}
		if isSiteAvailable() {
			ctx := buildOnce()
			if ctx != nil && cliCtx.Bool("watch") {
				watchCtx = ctx
				watchLoop()
				select {}
			}
		}
		return nil
	},
}

func buildOnce() *compile.Context {
	t := time.Now()
	printer.Print("=== start build ===")

	var (
		ctx *compile.Context
		err error
	)
	if ctx, err = compile.Read(); err != nil {
		printer.Error("read fail : %v", err)
		return nil
	}

	if err = compile.Compile(ctx); err != nil {
		printer.Error("compile fail : %v", err)
		return nil
	}

	if err = compile.Copy(ctx); err != nil {
		printer.Error("copy fail : %v", err)
		return nil
	}

	printer.Print("=== finish build ===")
	printer.Print("total in %s", time.Since(t))
	return ctx
}

func watchLoop() {
	printer.Print("--- watching...")
	w, err := watcher.New()
	if err != nil {
		printer.Error("watch error : %v", err)
		return
	}
	defer w.Close()

	ctx := watchCtx
	if err = w.Add(ctx.Config.PostDir, ctx.Config.PageDir, ctx.Config.ThemeDir, ctx.Config.LangDir, ctx.Meta.SrcFile); err != nil {
		printer.Error("watch error : %v", err)
		return
	}
	if err = w.Start(func() {
		buildOnce()
	}); err != nil {
		printer.Error("watch error : %v", err)
		return
	}
}
