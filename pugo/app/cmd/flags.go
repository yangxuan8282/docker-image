package cmd

import (
	"github.com/Unknwon/com"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/urfave/cli"
)

var commonFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "debug",
		Usage: "print all debug info",
	},
}

func isSiteAvailable() bool {
	printer.Logf("check meta.toml")
	if !com.IsFile("meta.toml") {
		printer.Error("error: meta.toml is not found.")
		printer.Print("you need create a new site here with 'pugo init'")
		return false
	}
	return true
}
