package cmd

import (
	"runtime"
	"time"

	"github.com/go-xiaohei/pugo/app/asset"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/vars"
	"github.com/urfave/cli"
)

// Version is command `version`
var Version = cli.Command{
	Name:  "version",
	Usage: "print version information",
	Action: func(cliCtx *cli.Context) error {
		printer.Info("PuGo version \t : %v", vars.Version)
		printer.Info("Release Date \t : %v", time.Now().Format("2006-01-02 15:04"))
		printer.Info("Golang version \t : %v", runtime.Version())
		printer.Info("Assets version \t : %v", asset.Date)
		return nil
	},
}
