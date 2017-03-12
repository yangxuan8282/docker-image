package page

import "github.com/go-xiaohei/pugo/app/model/author"

// FillAuthors fills authors to pages
func FillAuthors(pages []*Page, ag author.Group) {
	for _, p := range pages {
		p.author = ag.Get(p.AuthorName)
	}
}
