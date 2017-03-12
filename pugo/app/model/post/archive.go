package post

import (
	"fmt"
	"time"

	"github.com/go-xiaohei/pugo/app/model"
)

var (
	_ model.Node = (*ArchiveGroup)(nil)
)

const (
	// ArchiveYearly means archives in years
	ArchiveYearly = iota + 1
	// ArchiveMonthly means archives in months
	ArchiveMonthly
)

type (
	// Archive is post-list with post period segment
	Archive struct {
		Period int
		Posts  []*Post
	}
	// ArchiveGroup is archive-list with period way type
	ArchiveGroup struct {
		Type     int
		Archives []*Archive
	}
)

// URL returns this link of archive page
func (ag *ArchiveGroup) URL() string {
	return "archive.html"
}

// CreateTime returns now, no usage
func (ag *ArchiveGroup) CreateTime() time.Time {
	return time.Now()
}

// UpdateTime returns now, no usage
func (ag *ArchiveGroup) UpdateTime() time.Time {
	return time.Now()
}

// Time returns time period string for this archive
func (a *Archive) Time() string {
	if a.Period <= 9999 {
		return fmt.Sprintf("%d", a.Period)
	}
	return fmt.Sprintf("%d-%02d", a.Period/100, a.Period%100)
}

// MakeArchived makes posts in archive list
func MakeArchived(posts []*Post, t int) *ArchiveGroup {
	var (
		ag          = &ArchiveGroup{Type: t}
		lastArchive *Archive
	)
	for _, p := range posts {
		var period int
		if t == ArchiveYearly {
			period = p.created.Year()
		}
		if t == ArchiveMonthly {
			period = p.created.Year()*100 + int(p.created.Month())
		}
		if lastArchive == nil {
			lastArchive = &Archive{
				Period: period,
				Posts:  []*Post{p},
			}
		} else {
			if lastArchive.Period == period {
				lastArchive.Posts = append(lastArchive.Posts, p)
			} else {
				ag.Archives = append(ag.Archives, lastArchive)
				lastArchive = nil
			}
		}
	}
	if lastArchive != nil {
		ag.Archives = append(ag.Archives, lastArchive)
	}
	return ag
}
