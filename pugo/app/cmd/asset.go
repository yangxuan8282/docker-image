package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/helper/ziper"
	"github.com/urfave/cli"
)

var assetFlag = commonFlags

// Asset is 'asset' command
var Asset = cli.Command{
	Name:  "asset",
	Usage: "convert assets to go source code [for developer]",
	Flags: assetFlag,
	Action: func(cliCtx *cli.Context) error {
		if cliCtx.Bool("debug") {
			printer.EnableLogf = true
		}
		convertAsset()
		return nil
	},
}

type convertResult struct {
	fileCount  int
	fileSize   int64
	encodeSize int64
}

func (c *convertResult) ratio() string {
	return fmt.Sprintf("%.1f", float64(c.encodeSize)/float64(c.fileSize)*100)
}

func convertAsset() {
	var (
		buf  = bytes.NewBuffer(nil)
		file = "app/asset/asset.go"
		err  error
	)

	buf.WriteString("package asset \n\n")
	buf.WriteString(`var Date = "` + time.Now().Format("2006-01-02 15:04") + `"`)
	buf.WriteString("\n\n")
	buf.WriteString("var Data = make(map[string]string)\n\n")
	buf.WriteString("func init(){\n")

	r, err := convertAssetDir(buf, "post", "page", "theme", "lang", "media")
	if err != nil {
		printer.Error("write asset error : %v", err)
		return
	}
	r2, err := convertAssetFile(buf, "meta.toml")
	if err != nil {
		printer.Error("write asset error : %v", err)
		return
	}
	r.fileCount += r2.fileCount
	r.fileSize += r2.fileSize
	r.encodeSize += r2.encodeSize

	buf.WriteString("}\n")

	os.MkdirAll("app/asset", os.ModePerm)
	if err = ioutil.WriteFile(file, buf.Bytes(), os.ModePerm); err != nil {
		printer.Error("write asset error : %v", err)
		return
	}
	printer.Info("write assets %v", file)
	printer.Info("write ratio  %v", r.ratio()+"%")
	printer.Info("write size   %v KB", buf.Len()/1024)
}

func convertAssetDir(buf *bytes.Buffer, dirs ...string) (*convertResult, error) {
	res := &convertResult{}
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if path.Base(p) == ".DS_Store" {
				return nil
			}
			str, err := ziper.GzipFileBase64(p)
			if err != nil {
				return err
			}
			res.fileCount++
			res.fileSize += info.Size()
			res.encodeSize += int64(len(str))
			buf.WriteString(`Data["` + p + `"] = "` + str + `"`)
			buf.WriteString("\n")
			printer.Trace("convert file %v", p)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func convertAssetFile(buf *bytes.Buffer, file string) (*convertResult, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	str, err := ziper.GzipFileBase64(file)
	if err != nil {
		return nil, err
	}
	r := &convertResult{
		fileCount:  1,
		fileSize:   info.Size(),
		encodeSize: int64(len(str)),
	}
	buf.WriteString(`Data["` + file + `"] = "` + str + `"`)
	buf.WriteString("\n")
	printer.Trace("convert file %v", file)
	return r, nil
}
