package post

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-xiaohei/pugo/app/model"
)

var (
	_ model.Node = (*Tag)(nil)
)

// Tag is an object for a post tag
type Tag struct {
	Name    string
	created time.Time
	updated time.Time
}

// URL returns tag posts html
func (t *Tag) URL() string {
	return fmt.Sprintf("tags/%s.html", t.Name)
}

// CreateTime returns the first post's created time with this tag
func (t *Tag) CreateTime() time.Time {
	return t.created
}

// UpdateTime returns the last post's updated time with this tag
func (t *Tag) UpdateTime() time.Time {
	return t.updated
}

// MakeTags makes post-list with its tag
func MakeTags(posts []*Post) map[string]*TagPosts {
	m := make(map[string]*TagPosts)
	for _, p := range posts {
		tagNames := p.TagStrings
		if len(tagNames) == 0 {
			continue
		}
		for _, tag := range tagNames {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			t := &Tag{
				Name: tag,
			}
			p.tags = append(p.tags, t)
			tagKey := strings.ToLower(tag)
			if m[tagKey] == nil {
				m[tagKey] = &TagPosts{
					Tag: t,
				}
			}
			m[tagKey].Posts = append(m[tagKey].Posts, p)
		}
	}
	for _, tp := range m {
		tp.Tag.created = tp.Posts[len(tp.Posts)-1].created
		tp.Tag.updated = tp.Posts[0].updated
	}
	return m
}
