package post

import (
	"fmt"
	"time"

	"github.com/go-xiaohei/pugo/app/helper/pager"
	"github.com/go-xiaohei/pugo/app/model"
	"github.com/go-xiaohei/pugo/app/model/author"
)

var (
	_ model.Node = (*PagedPosts)(nil)
	_ model.Node = (*TagPosts)(nil)
)

type (
	// Posts is list of post
	Posts []*Post
	// DatePosts is list of post. it can be sorted by created time
	DatePosts Posts
)

func (dp DatePosts) Len() int { return len(dp) }

func (dp DatePosts) Swap(i, j int) { dp[i], dp[j] = dp[j], dp[i] }

func (dp DatePosts) Less(i, j int) bool { return dp[i].created.Unix() > dp[j].created.Unix() }

type (
	// PagedPosts is list of post with it's pager
	PagedPosts struct {
		Posts []*Post
		Pager *pager.Pager
	}
	// TagPosts is list of post with its tag
	TagPosts struct {
		Posts []*Post
		Tag   *Tag
	}
)

// URL returns current url of this paged post-list
func (pp *PagedPosts) URL() string {
	return fmt.Sprintf("posts/%d.html", pp.Pager.Current)
}

// CreateTime returns the first post's created time in paged post-list
func (pp *PagedPosts) CreateTime() time.Time {
	return pp.Posts[len(pp.Posts)-1].created
}

// UpdateTime returns the last post's updated time in paged post-list
func (pp *PagedPosts) UpdateTime() time.Time {
	return pp.Posts[0].updated
}

// URL returns current url of this tagged post-list
func (tp *TagPosts) URL() string {
	return tp.Tag.URL()
}

// CreateTime returns the first post's created time in tagged post-list
func (tp *TagPosts) CreateTime() time.Time {
	return tp.Tag.created
}

// UpdateTime returns the last post's updated time in tagged post-list
func (tp *TagPosts) UpdateTime() time.Time {
	return tp.Tag.updated
}

// MakePaged makes posts to paged post-list with size of per page
func MakePaged(posts []*Post, size int) []*PagedPosts {
	var (
		p          = pager.NewCursor(4, len(posts))
		i          = 1
		pagerPosts []*PagedPosts
	)
	for {
		pg := p.Page(i)
		if pg == nil {
			break
		}
		pp := &PagedPosts{
			Pager: pg,
			Posts: posts[pg.Begin:pg.End],
		}
		pagerPosts = append(pagerPosts, pp)
		i++
	}
	return pagerPosts
}

// FillAuthors fills authors to posts
func FillAuthors(posts []*Post, ag author.Group) {
	for _, p := range posts {
		p.author = ag.Get(p.AuthorName)
	}
}
