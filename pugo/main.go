package main

import (
	"os"
	"time"

	"github.com/go-xiaohei/pugo/app/cmd"
	"github.com/go-xiaohei/pugo/app/vars"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "PuGo"
	app.Usage = "a Simple Static Site Generator"
	app.UsageText = app.Usage
	app.Commands = []cli.Command{
		cmd.Init,
		cmd.Build,
		cmd.Server,
		cmd.Asset,
		cmd.Version,
	}
	app.Version = vars.Version
	app.Compiled = time.Now()
	app.HideVersion = true
	// app.HideHelp = true
	app.Run(os.Args)
}
