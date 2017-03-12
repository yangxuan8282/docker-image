package compile

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/go-xiaohei/pugo/app/helper/i18n"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/go-xiaohei/pugo/app/model/author"
	"github.com/go-xiaohei/pugo/app/model/base"
	"github.com/go-xiaohei/pugo/app/model/node"
	"github.com/go-xiaohei/pugo/app/model/page"
	"github.com/go-xiaohei/pugo/app/model/post"
	"github.com/go-xiaohei/pugo/app/model/theme"
	"github.com/go-xiaohei/pugo/app/model/third"
	"github.com/go-xiaohei/pugo/app/vars"
)

var (
	// ErrContextThemeMissing means theme is missing when compiling with context
	ErrContextThemeMissing = errors.New("context missing theme when compiling")
)

type (
	// Context is a context in building once
	Context struct {
		Config    *base.Config     `toml:"config"`
		Meta      *base.Meta       `toml:"meta"`
		Authors   []*author.Author `toml:"author"`
		Navs      base.NavGroup    `toml:"nav"`
		Comment   *third.Comment   `toml:"comment"`
		Analytics *third.Analytics `toml:"analytics"`

		posts     []*post.Post
		pages     []*page.Page
		tags      []*post.Tag
		theme     *theme.Theme
		i18nGroup *i18n.Group
		drafts    []interface{}

		pagerPosts []*post.PagedPosts
		indexPosts *post.PagedPosts
		tagPosts   map[string]*post.TagPosts
		archives   *post.ArchiveGroup
		tree       *node.Node
	}
)

// DstDir returns the real destination directory,
// it combines Meta.Root and Config.OutputDir
func (ctx *Context) DstDir() string {
	u, _ := ctx.Meta.RootURL()
	return path.Join(ctx.Config.OutputDir, u.Path)
}

// ViewData returns common view data from context
func (ctx *Context) ViewData() map[string]interface{} {
	u, _ := ctx.Meta.RootURL()
	return map[string]interface{}{
		"Meta":      ctx.Meta,
		"I18n":      ctx.i18nGroup.Get(ctx.Meta.Lang),
		"I18nList":  ctx.i18nGroup.List,
		"I18nItem":  ctx.i18nGroup.GetItem(ctx.Meta.Lang),
		"Nav":       ctx.Navs,
		"Hover":     "",
		"Now":       time.Now(),
		"Version":   vars.Version,
		"Title":     ctx.Meta.Title + " - " + ctx.Meta.Subtitle,
		"Keyword":   ctx.Meta.Keyword,
		"Desc":      ctx.Meta.Description,
		"Comment":   ctx.Comment,
		"Analytics": ctx.Analytics,
		"Base":      strings.TrimSuffix(u.Path, "/"),
		"Tree":      ctx.tree,
		"Slug":      "",
	}
}

func (ctx *Context) compile(tpl string, viewData map[string]interface{}) ([]byte, error) {
	if ctx.theme == nil {
		return nil, ErrContextThemeMissing
	}
	buf := bytes.NewBuffer(nil)
	data := ctx.ViewData()
	for k, v := range viewData {
		data[k] = v
	}
	if err := ctx.theme.Execute(buf, tpl, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (ctx *Context) buildTree() {
	u, _ := ctx.Meta.RootURL()
	ctx.tree = node.NewTree(u.Path)
	ctx.tree.FillCommonPages()
	ctx.tree.FillPosts(ctx.posts)
	ctx.tree.FillPages(ctx.pages)
	ctx.tree.FillPagedPosts(ctx.pagerPosts)
	ctx.tree.FillTagPosts(ctx.tagPosts)
	ctx.tree.SortChildren()
	// ctx.tree.Print("")
}

func (ctx *Context) readAssemble() error {
	post.FillAuthors(ctx.posts, author.Group(ctx.Authors))
	page.FillAuthors(ctx.pages, author.Group(ctx.Authors))

	ctx.pagerPosts = post.MakePaged(ctx.posts, ctx.Config.PostPageSize)
	printer.Logf("assemble page-posts %d", len(ctx.pagerPosts))
	if len(ctx.pagerPosts) > 0 {
		ctx.indexPosts = ctx.pagerPosts[0]
	}

	ctx.tagPosts = post.MakeTags(ctx.posts)
	printer.Logf("assemble tag-posts %d", len(ctx.tagPosts))

	ctx.archives = post.MakeArchived(ctx.posts, post.ArchiveYearly)
	printer.Logf("assemble archives %d", len(ctx.archives.Archives))

	ctx.buildTree()
	printer.Logf("assemble tree %d nodes", ctx.tree.Len())
	return nil
}

func (ctx *Context) prepareCompileTask() []compileTask {
	var tasks []compileTask
	for _, p := range ctx.posts {
		tasks = append(tasks, compileTask{
			TplFile: "post.html",
			DstFile: path.Join(ctx.DstDir(), p.URL()),
			ViewData: map[string]interface{}{
				"Post":  p,
				"Title": p.Title + " - " + ctx.Meta.Title,
				"Slug":  p.URL(),
			},
			SrcFile: p.SrcFile(),
		})
	}
	for _, p := range ctx.pages {
		if p.IsNode {
			continue
		}
		var lang = p.Lang
		if lang == "" {
			lang = ctx.Meta.Lang
		}
		task := compileTask{
			TplFile: "page.html",
			DstFile: path.Join(ctx.DstDir(), p.URL()),
			ViewData: map[string]interface{}{
				"Page":     p,
				"Hover":    p.Hover,
				"Title":    p.Title + " - " + ctx.Meta.Title,
				"Slug":     p.URL(),
				"I18n":     ctx.i18nGroup.Get(lang),
				"I18nItem": ctx.i18nGroup.GetItem(lang),
			},
			SrcFile: p.SrcFile(),
		}
		if p.Template != "" {
			task.TplFile = p.Template
		}
		tasks = append(tasks, task)
	}
	for _, pp := range ctx.pagerPosts {
		task := compileTask{
			TplFile: "posts.html",
			DstFile: path.Join(ctx.DstDir(), pp.URL()),
			ViewData: map[string]interface{}{
				"Pager": pp.Pager,
				"Posts": pp.Posts,
				"Title": fmt.Sprintf("Page %d - %s", pp.Pager.Current, ctx.Meta.Title),
				"Slug":  pp.URL(),
			},
			SrcFile: fmt.Sprintf("post-list-%d", pp.Pager.Current),
		}
		tasks = append(tasks, task)
	}
	indexPageTask := compileTask{
		TplFile: "posts.html",
		DstFile: path.Join(ctx.DstDir(), "index.html"),
		ViewData: map[string]interface{}{
			"Pager": ctx.indexPosts.Pager,
			"Posts": ctx.indexPosts.Posts,
			"Hover": "index",
			"Slug":  "index.html",
		},
		SrcFile: "index",
	}
	if ctx.theme.Template("index.html") != nil {
		indexPageTask.TplFile = "index.html"
	}
	tasks = append(tasks, indexPageTask)
	for _, tp := range ctx.tagPosts {
		task := compileTask{
			TplFile: "posts.html",
			DstFile: path.Join(ctx.DstDir(), tp.URL()),
			ViewData: map[string]interface{}{
				"Tag":   tp.Tag,
				"Posts": tp.Posts,
				"Title": fmt.Sprintf("Tag %d - %s", tp.Tag.Name, ctx.Meta.Title),
				"Slug":  tp.URL(),
			},
			SrcFile: fmt.Sprintf("post-tag-%s", tp.Tag.Name),
		}
		tasks = append(tasks, task)
	}
	archiveTask := compileTask{
		TplFile: "archive.html",
		DstFile: path.Join(ctx.DstDir(), ctx.archives.URL()),
		ViewData: map[string]interface{}{
			"Archives": ctx.archives.Archives,
			"Hover":    "archive",
			"Title":    "Archive - " + ctx.Meta.Title,
			"Slug":     ctx.archives.URL(),
		},
		SrcFile: "post-archive",
	}
	tasks = append(tasks, archiveTask)
	return tasks
}
