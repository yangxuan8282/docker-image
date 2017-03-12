package compile

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/go-xiaohei/pugo/app/helper/printer"
)

// Copy copies static files to destination directory
func Copy(ctx *Context) error {
	var (
		copySuccessCounter int
		copyFailCounter    int
	)
	// copy media
	filepath.Walk(ctx.Config.MediaDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		printer.Logf("copy %s", p)
		dstFile := path.Join(ctx.DstDir(), p)
		os.MkdirAll(path.Dir(dstFile), os.ModePerm)
		if err = com.Copy(p, dstFile); err != nil {
			printer.Error("copy %s error : %v", p, err)
			copyFailCounter++
			return nil
		}
		copySuccessCounter++
		return nil
	})

	// copy theme static
	filepath.Walk(ctx.theme.StaticDir(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if path.Ext(p) == ".DS_Store" {
			return nil
		}
		printer.Logf("copy %s", p)
		rel, _ := filepath.Rel(ctx.theme.StaticDir(), p)
		dstFile := path.Join(ctx.DstDir(), rel)
		os.MkdirAll(path.Dir(dstFile), os.ModePerm)
		if err = com.Copy(p, dstFile); err != nil {
			printer.Error("copy %s error : %v", p, err)
			copyFailCounter++
			return nil
		}
		copySuccessCounter++
		return nil
	})
	printer.Info("copied files \t: %v", copySuccessCounter)
	if copyFailCounter > 0 {
		printer.Warn("copied fails \t: %v", copyFailCounter)
	}
	return nil
}
