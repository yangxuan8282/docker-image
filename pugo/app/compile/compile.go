package compile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/gorilla/feeds"
)

type compileTask struct {
	TplFile  string
	ViewData map[string]interface{}
	DstFile  string
	SrcFile  string
}

// Compile compiles context data to destination files
func Compile(ctx *Context) error {
	printer.Print("--- compiling... ")
	printer.Logf("compile to %s", ctx.DstDir())

	var (
		tasks              = ctx.prepareCompileTask()
		taskSuccessCounter int
		taskFailCounter    int
	)
	printer.Logf("compile tasks %d", len(tasks))

	for _, task := range tasks {
		data, err := ctx.compile(task.TplFile, task.ViewData)
		if err != nil {
			printer.Error("compile %v fail : %v", task.SrcFile, err)
			taskFailCounter++
		} else {
			os.MkdirAll(path.Dir(task.DstFile), os.ModePerm)
			if err = ioutil.WriteFile(task.DstFile, data, os.ModePerm); err != nil {
				printer.Error("compile %v fail : %v", task.SrcFile, err)
				taskFailCounter++
			}
		}
		if err == nil {
			taskSuccessCounter++
		}
	}

	printer.Logf("compile rss")
	if err := compileRSS(ctx); err != nil {
		printer.Error("compile rss fail : %v", err)
		taskFailCounter++
	} else {
		taskSuccessCounter++
	}

	printer.Logf("compile sitemap")
	if err := compileSitemap(ctx); err != nil {
		printer.Error("compile sitemap fail : %v", err)
		taskFailCounter++
	} else {
		taskSuccessCounter++
	}

	printer.Info("created pages \t: %v", taskSuccessCounter)
	if taskFailCounter > 0 {
		printer.Warn("failed pages \t: %v", taskFailCounter)
	}
	return nil
}

func compileRSS(ctx *Context) error {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       ctx.Meta.Title,
		Link:        &feeds.Link{Href: ctx.Meta.Root},
		Description: ctx.Meta.Description,
		Created:     now,
	}
	if len(ctx.Authors) > 0 {
		feed.Author = &feeds.Author{
			Name:  ctx.Authors[0].Nick,
			Email: ctx.Authors[0].Email,
		}
	}
	var item *feeds.Item
	for _, p := range ctx.posts {
		u, _ := ctx.Meta.RootURL()
		u.Path = path.Join(u.Path, p.URL())
		item = &feeds.Item{
			Title:       p.Title,
			Link:        &feeds.Link{Href: u.String()},
			Description: string(p.Content()),
			Created:     p.CreateTime(),
			Updated:     p.UpdateTime(),
		}
		if p.Author() != nil {
			item.Author = &feeds.Author{
				Name:  p.Author().Nick,
				Email: p.Author().Email,
			}
		}
		feed.Items = append(feed.Items, item)
	}

	dstFile := path.Join(ctx.DstDir(), "feed.xml")
	os.MkdirAll(filepath.Dir(dstFile), os.ModePerm)
	f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return feed.WriteRss(f)
}

func compileSitemap(ctx *Context) error {
	now := time.Now()
	var (
		buf bytes.Buffer
		u   *url.URL
	)
	buf.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	buf.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	buf.WriteString("<url>")
	fmt.Fprintf(&buf, "<loc>%s</loc>", ctx.Meta.Root)
	fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", now.Format(time.RFC3339))
	buf.WriteString("<changefreq>daily</changefreq>")
	buf.WriteString("<priority>1.0</priority>")
	buf.WriteString("</url>")

	for _, p := range ctx.posts {
		u, _ = ctx.Meta.RootURL()
		u.Path = path.Join(u.Path, p.URL())
		buf.WriteString("<url>")
		fmt.Fprintf(&buf, "<loc>%s</loc>", u.String())
		fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", p.CreateTime().Format(time.RFC3339))
		buf.WriteString("<changefreq>daily</changefreq>")
		buf.WriteString("<priority>0.6</priority>")
		buf.WriteString("</url>")
	}

	for _, p := range ctx.pages {
		u, _ = ctx.Meta.RootURL()
		u.Path = path.Join(u.Path, p.URL())
		buf.WriteString("<url>")
		fmt.Fprintf(&buf, "<loc>%s</loc>", u.String())
		fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", p.CreateTime().Format(time.RFC3339))
		buf.WriteString("<changefreq>weekly</changefreq>")
		buf.WriteString("<priority>0.5</priority>")
		buf.WriteString("</url>")
	}

	u, _ = ctx.Meta.RootURL()
	u.Path = path.Join(u.Path, "archive.html")
	buf.WriteString("<url>")
	fmt.Fprintf(&buf, "<loc>%s</loc>", u.String())
	fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", now.Format(time.RFC3339))
	buf.WriteString("<changefreq>daily</changefreq>")
	buf.WriteString("<priority>0.6</priority>")
	buf.WriteString("</url>")

	for _, pp := range ctx.pagerPosts {
		u, _ = ctx.Meta.RootURL()
		u.Path = path.Join(u.Path, fmt.Sprintf("posts/%d.html", pp.Pager.Current))
		buf.WriteString("<url>")
		fmt.Fprintf(&buf, "<loc>%s</loc>", u.String())
		fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", now.Format(time.RFC3339))
		buf.WriteString("<changefreq>daily</changefreq>")
		buf.WriteString("<priority>0.6</priority>")
		buf.WriteString("</url>")
	}

	for _, tp := range ctx.tagPosts {
		u, _ = ctx.Meta.RootURL()
		u.Path = path.Join(u.Path, tp.Tag.URL())
		buf.WriteString("<url>")
		fmt.Fprintf(&buf, "<loc>%s</loc>", u.String())
		fmt.Fprintf(&buf, "<lastmod>%s</lastmod>", now.Format(time.RFC3339))
		buf.WriteString("<changefreq>weekly</changefreq>")
		buf.WriteString("<priority>0.5</priority>")
		buf.WriteString("</url>")
	}

	buf.WriteString("</urlset>")
	dstFile := path.Join(ctx.DstDir(), "sitemap.xml")
	os.MkdirAll(path.Dir(dstFile), os.ModePerm)
	if err := ioutil.WriteFile(dstFile, buf.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}
