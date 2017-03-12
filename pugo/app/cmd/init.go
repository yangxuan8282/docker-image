package cmd

import (
	"os"

	"github.com/go-xiaohei/pugo/app/asset"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/helper/ziper"
	"github.com/urfave/cli"
)

var initFlags = []cli.Flag{}

// Init is 'init' command
var Init = cli.Command{
	Name:  "init",
	Usage: "create new site, content and theme for your static site",
	Flags: initFlags,
	Action: func(ctx *cli.Context) error {
		if isSiteAvailable() {
			printer.Error("error: meta.toml is found.")
			printer.Trace("you had created a site here.")
			return nil
		}
		printer.Trace("read assets %v", len(asset.Data))
		for file, data := range asset.Data {
			if err := ziper.UnGzipFileBase64(data, file); err != nil {
				printer.Error("extract asset %v error : %v", file, err)
				return nil
			}
			printer.Trace("extract asset %v", file)
		}
		wd, _ := os.Getwd()
		printer.Info("Congratulations! PuGo default site is created in %v", wd)
		return nil
	},
}
