package compile

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"
	"github.com/go-xiaohei/pugo/app/helper/i18n"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/model/page"
	"github.com/go-xiaohei/pugo/app/model/post"
	"github.com/go-xiaohei/pugo/app/model/theme"
	"github.com/go-xiaohei/pugo/app/vars"
)

var (
	// ErrReadMetaInvalid means invalid meta file
	ErrReadMetaInvalid = errors.New("read invalid meta file")
	// ErrMetaLangInvalid means the language in meta file is not found in language directory
	ErrMetaLangInvalid = errors.New("meta's language is invalid")
	// ErrReadMetaFail means it can not load correct meta file
	ErrReadMetaFail = errors.New("read meta fail")
	// ErrThemeDirInvalid means theme dir is empty
	ErrThemeDirInvalid = errors.New("theme dir is invalid")
	// ErrDestDirInvalid means dest dir is invalid
	ErrDestDirInvalid = errors.New("dest dir is invalid")
)

// Payload loads meta for basic Context
func Payload() (*Context, error) {
	var (
		err error
		ctx = new(Context)
	)
	if err = readMeta(ctx); err != nil {
		return nil, err
	}
	if _, err = ctx.Meta.RootURL(); err != nil {
		return nil, err
	}
	if ctx.Config.OutputDir == "" {
		return nil, ErrDestDirInvalid
	}
	return ctx, err
}

// Read reads all contents and parses to model objects
func Read() (*Context, error) {
	ctx, err := Payload()
	if err != nil {
		return nil, err
	}
	printer.Print("--- reading...")
	printer.Trace("language \t: %v", ctx.Meta.Lang)
	printer.Trace("Authors \t: %v", len(ctx.Authors))
	printer.Trace("Navigators \t: %v", len(ctx.Navs))

	if err = readPosts(ctx); err != nil {
		return nil, err
	}
	if err = readPages(ctx); err != nil {
		return nil, err
	}
	if err = readLang(ctx); err != nil {
		return nil, err
	}
	if err = readTheme(ctx); err != nil {
		return nil, err
	}
	if err = ctx.readAssemble(); err != nil {
		return nil, err
	}
	printer.Trace("post tags \t: %v", len(ctx.tagPosts))

	return ctx, nil
}

func readMeta(ctx *Context) error {
	var (
		fileBytes []byte
		err       error
		isFind    bool
	)

	for _, file := range vars.MetaFiles {
		printer.Logf("try read %s", file)
		if !com.IsFile(file) {
			continue
		}
		fileBytes, err = ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		printer.Logf("read %s", file)
		if err = toml.Unmarshal(fileBytes, ctx); err != nil {
			return err
		}
		if ctx.Meta == nil || ctx.Config == nil {
			printer.Error("error: %v", ErrReadMetaInvalid)
			printer.Print("meta file need [meta] and [base] sections")
			continue
		}
		ctx.Meta.SrcFile = file
		isFind = true
		break
	}
	if !isFind {
		return ErrReadMetaFail
	}

	if len(ctx.Authors) > 0 {
		for _, author := range ctx.Authors {
			if err = author.Format(); err != nil {
				printer.Error("author %v error: %v", author.Name, err)
			}
		}
		ctx.Authors[0].IsOwner = true
	}
	if len(ctx.Navs) > 0 {
		if err = ctx.Navs.Format(); err != nil {
			printer.Error("nav %v error: %v")
		}
	}
	return err
}

func readPosts(ctx *Context) error {
	if ctx.Config.PostDir == "" {
		return nil
	}
	var (
		posts  []*post.Post
		drafts int
	)
	err := filepath.Walk(ctx.Config.PostDir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if path.Ext(info.Name()) != ".md" {
			return nil
		}
		printer.Logf("read %s", file)
		p, err := post.NewFromFile(file)
		if err != nil {
			printer.Error("parse file %v error : %v", file, err)
			return nil
		}
		if !p.IsDraft {
			posts = append(posts, p)
		} else {
			ctx.drafts = append(ctx.drafts, p)
			drafts++
		}
		return nil
	})
	if len(posts) > 1 {
		sort.Sort(post.DatePosts(posts))
	}
	ctx.posts = posts
	printer.Trace("post files \t: %v", len(posts))
	if drafts > 0 {
		printer.Warn("post drafts \t: %v", drafts)
	}
	return err
}

func readPages(ctx *Context) error {
	if ctx.Config.PageDir == "" {
		return nil
	}
	var (
		pages  []*page.Page
		drafts int
	)
	err := filepath.Walk(ctx.Config.PageDir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if path.Ext(info.Name()) != ".md" {
			return nil
		}
		printer.Logf("read %s", file)
		rel, _ := filepath.Rel(ctx.Config.PageDir, file)
		rel = strings.TrimSuffix(rel, path.Ext(rel))
		p, err := page.NewFromFile(file, rel)
		if err != nil {
			printer.Error("parse file %v error : %v", file, err)
			return nil
		}
		if !p.IsDraft {
			pages = append(pages, p)
		} else {
			ctx.drafts = append(ctx.drafts, p)
			drafts++
		}
		return nil
	})
	ctx.pages = pages
	printer.Trace("page files \t: %v", len(pages))
	if drafts > 0 {
		printer.Trace("page drafts \t: %v", drafts)
	}
	return err
}

func readLang(ctx *Context) error {
	if ctx.Meta.Lang == "" {
		ctx.i18nGroup = i18n.NewGroup()
		return nil
	}
	if ctx.Config.LangDir == "" {
		return ErrMetaLangInvalid
	}
	i18nGroup := i18n.NewGroup()

	i18nListFile := filepath.Join(ctx.Config.LangDir, vars.I18nListFile)
	if !com.IsFile(i18nListFile) {
		return i18n.ErrI18nGroupListMissing
	}
	data, err := ioutil.ReadFile(i18nListFile)
	if err != nil {
		return err
	}
	printer.Logf("read %s", i18nListFile)
	if err = toml.Unmarshal(data, i18nGroup); err != nil {
		return err
	}
	if err = i18nGroup.Validate(); err != nil {
		return err
	}
	for _, item := range i18nGroup.List {
		file := filepath.Join(ctx.Config.LangDir, item.File)
		if !com.IsFile(file) {
			printer.Warn("language file %s is missing", item.File)
			continue
		}
		in, err := i18n.NewFromFile(file)
		if err != nil {
			return err
		}
		i18nGroup.Set(in)
		printer.Logf("read %s", file)
	}
	ctx.i18nGroup = i18nGroup
	printer.Trace("language files \t: %v", strings.Join(ctx.i18nGroup.Names(), ","))
	if ctx.i18nGroup.Get(ctx.Meta.Lang) == nil {
		return ErrMetaLangInvalid
	}
	return err
}

func readTheme(ctx *Context) error {
	if ctx.Config.ThemeDir == "" {
		return ErrThemeDirInvalid
	}
	t := theme.New(ctx.Config.ThemeDir)
	if err := t.Validate(); err != nil {
		return err
	}
	t.Func("FullURL", func(p string) string {
		if strings.Contains(p, "//") {
			return p
		}
		u, _ := ctx.Meta.RootURL()
		p = path.Join(u.Path, p)
		if !strings.HasPrefix(p, "/") {
			return "/" + p
		}
		return p
	})
	t.Func("FullTrURL", func(in *i18n.I18n, p string) string {
		if strings.Contains(p, "//") {
			return p
		}
		u, _ := ctx.Meta.RootURL()
		lang := ""
		if in != nil {
			lang = in.Lang
		}
		p = path.Join(u.Path, lang, p)
		if !strings.HasPrefix(p, "/") {
			return "/" + p
		}
		return p
	})
	if err := t.Load(); err != nil {
		return err
	}
	ctx.theme = t
	printer.Trace("theme files \t: %v", t.Len())
	return nil
}
