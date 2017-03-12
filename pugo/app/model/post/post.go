// Package post delares the types of post, tag and post-list
package post

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-xiaohei/pugo/app/helper/markdown"
	"github.com/go-xiaohei/pugo/app/model"
	"github.com/go-xiaohei/pugo/app/model/author"
	"github.com/go-xiaohei/pugo/app/model/index"
	"github.com/go-xiaohei/pugo/app/vars"
)

var (
	_ model.Content = (*Post)(nil)
	_ model.Index   = (*Post)(nil)
)

var (
	postContentBreak = []byte("<!--more-->")
)

var (
	// ErrPostFrontMetaFail means it can't detect front-matter block in post bytes
	ErrPostFrontMetaFail = errors.New("detect front-matter fail")
	// ErrPostFrontMetaTypeUnknown means it can't parse front-matter block with known types
	ErrPostFrontMetaTypeUnknown = errors.New("can't detect front-matter's format")
	// ErrPostFrontMetaTimeError means wrong time format in front-matter block
	ErrPostFrontMetaTimeError = errors.New("time format error in front-matter")
)

type (
	// Post is an object for one post content
	Post struct {
		Title      string   `toml:"title"`
		Slug       string   `toml:"slug"`
		Desc       string   `toml:"desc"`
		Created    string   `toml:"date"`
		Updated    string   `toml:"update_date"`
		AuthorName string   `toml:"author"`
		Thumb      string   `toml:"thumb"`
		TagStrings []string `toml:"tags"`
		IsDraft    bool     `toml:"draft"`

		index  []*index.Index
		tags   []*Tag
		author *author.Author

		briefBytes   []byte
		contentBytes []byte
		srcBytes     []byte
		srcFile      string
		created      time.Time
		updated      time.Time

		frontMetaBytes []byte
		frontMetaType  int
	}
)

func (p *Post) detectFrontMeta() error {
	dataSlice := bytes.SplitN(p.srcBytes, vars.FrontMetaBreak, 3)
	if len(dataSlice) != 3 {
		return ErrPostFrontMetaFail
	}
	frontBytes := bytes.TrimSpace(dataSlice[1])
	for t, prefix := range vars.FrontMetaTypes {
		if bytes.HasPrefix(frontBytes, prefix) {
			frontBytes = bytes.TrimPrefix(frontBytes, prefix)
			p.frontMetaBytes = frontBytes
			p.frontMetaType = t
			p.contentBytes = bytes.TrimSpace(dataSlice[2])
			return nil
		}
	}
	return ErrPostFrontMetaTypeUnknown
}

func (p *Post) parseFrontMeta() error {
	var err error
	if err = toml.Unmarshal(p.frontMetaBytes, p); err != nil {
		return err
	}
	if err = p.formatTime(); err != nil {
		return err
	}
	return nil
}

func (p *Post) formatTime() error {
	var err error
	if p.Created == "" {
		p.getCreateTime()
	} else {
		for _, layout := range vars.TimeFormatLayout {
			p.created, err = time.Parse(layout, p.Created)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
	}
	if p.Updated == "" {
		p.Updated = p.Created
		p.updated = p.created
	} else {
		for _, layout := range vars.TimeFormatLayout {
			p.updated, err = time.Parse(layout, p.Updated)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Post) getBrief() {
	dataSlice := bytes.SplitN(p.contentBytes, postContentBreak, 2)
	if len(dataSlice) < 2 || len(dataSlice[1]) == 0 {
		p.briefBytes = p.contentBytes
		return
	}
	p.briefBytes = dataSlice[0]
}

func (p *Post) getCreateTime() {
	if p.srcFile != "" {
		if info, _ := os.Stat(p.srcFile); info != nil {
			p.created = info.ModTime()
			p.Created = p.created.Format("2006-01-02 15:04")
		}
	}
}

func (p *Post) render() {
	p.briefBytes = markdown.Render(p.briefBytes)
	p.contentBytes = markdown.Render(p.contentBytes)
}

func (p *Post) getIndex() {
	p.index = index.New(p.contentBytes)
}

// New parses bytes to a *Post
func New(dataBytes []byte, file string) (*Post, error) {
	var (
		err error
		p   = &Post{
			srcBytes: dataBytes,
			srcFile:  file,
		}
	)
	if err = p.detectFrontMeta(); err != nil {
		return nil, err
	}
	if err = p.parseFrontMeta(); err != nil {
		return nil, err
	}
	if p.IsDraft {
		return p, nil
	}
	p.getBrief()
	p.render()
	p.getIndex()
	return p, nil
}

// NewFromFile parses file to a *Post
func NewFromFile(file string) (*Post, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	p, err := New(data, file)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Content returns post's content
func (p *Post) Content() []byte {
	return p.contentBytes
}

// ContentHTML returns post's content html type
func (p *Post) ContentHTML() template.HTML {
	return template.HTML(p.contentBytes)
}

// ContentLength return post's content length
func (p *Post) ContentLength() int {
	return len(p.contentBytes)
}

// Brief returns post's brief content
func (p *Post) Brief() []byte {
	return p.briefBytes
}

// BriefHTML returns post's brief content html type
func (p *Post) BriefHTML() template.HTML {
	return template.HTML(p.briefBytes)
}

// CreateTime returns post's created time
func (p *Post) CreateTime() time.Time {
	return p.created
}

// UpdateTime returns post's updated time
func (p *Post) UpdateTime() time.Time {
	return p.updated
}

// DstFile returns rendered destination filepath
func (p *Post) DstFile() string {
	if p.IsDraft {
		return ""
	}
	return fmt.Sprintf("%s/%s.html", p.created.Format("2006/1/2"), p.Slug)
}

// SrcFile returns source filepath
func (p *Post) SrcFile() string {
	return p.srcFile
}

// URL returns site link for this post
func (p *Post) URL() string {
	if p.IsDraft {
		return ""
	}
	return fmt.Sprintf("%s/%s.html", p.created.Format("2006/1/2"), p.Slug)
}

// Index returns content index for post
func (p *Post) Index() []*index.Index {
	return p.index
}

// Tags returns tags of this post
func (p *Post) Tags() []*Tag {
	return p.tags
}

// Author gets the author pf this post
func (p *Post) Author() *author.Author {
	return p.author
}
