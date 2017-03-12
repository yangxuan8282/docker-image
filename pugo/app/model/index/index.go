package index

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Index is index of post
type Index struct {
	Level    int
	Title    string
	Anchor   string
	Children []*Index
	Link     string
	Parent   *Index
}

// Print prints post indexs friendly
func (p *Index) Print() {
	fmt.Println(strings.Repeat("#", p.Level), p)
	for _, c := range p.Children {
		c.Print()
	}
}

// New returns indexes with data bytes
func New(data []byte) []*Index {
	return NewWithReader(bytes.NewReader(data))
}

// NewWithReader returns indexes with a reader
func NewWithReader(r io.Reader) []*Index {
	var (
		z = html.NewTokenizer(r)

		currentLevel    int
		currentText     string
		currentLinkText string
		currentAnchor   string
		nodeDeep        int

		indexs []*Index
	)
	for {
		token := z.Next()
		if token == html.ErrorToken {
			break
		}
		if token == html.EndTagToken {
			if nodeDeep == 1 && currentLevel > 0 {
				indexs = append(indexs, &Index{
					Level:  currentLevel,
					Title:  currentText,
					Link:   currentLinkText,
					Anchor: currentAnchor,
				})
				currentLevel = 0
				currentText = ""
				currentLinkText = ""
				currentAnchor = ""
				nodeDeep--
			}
			continue
		}
		if token == html.StartTagToken {
			name, hasAttr := z.TagName()
			lv := parseIndexLevel(name)

			if lv > 0 {
				currentLevel = lv
				if hasAttr {
					for {
						k, v, isMore := z.TagAttr()
						if bytes.Equal(k, []byte("id")) {
							currentAnchor = string(v)
						}
						if !isMore {
							break
						}
					}
				}
				nodeDeep++
			}

			if currentLevel > 0 && string(name) == "a" {
				if hasAttr {
					for {
						k, v, isMore := z.TagAttr()
						if bytes.Equal(k, []byte("href")) {
							currentLinkText = string(v)
						}
						if !isMore {
							break
						}
					}
				}
			}
		}
		if token == html.TextToken && currentLevel > 0 {
			currentText += string(z.Text())
		}
	}
	indexs = assembleIndex(indexs)
	return indexs
}

func assembleIndex(indexList []*Index) []*Index {
	var (
		list    []*Index
		lastIdx int
		lastN   *Index
	)
	for i, n := range indexList {
		if i == 0 {
			list = append(list, n)
			lastIdx = 0
			continue
		}
		lastN = list[lastIdx]
		if lastN.Level < n.Level {
			n.Parent = lastN
			lastN.Children = append(lastN.Children, n)
		} else {
			list = append(list, n)
			lastIdx++
		}
	}
	for _, n := range list {
		if len(n.Children) > 1 {
			n.Children = assembleIndex(n.Children)
		}
	}
	return list
}

func parseIndexLevel(name []byte) int {
	if bytes.Equal(name, []byte("h1")) {
		return 1
	}
	if bytes.Equal(name, []byte("h2")) {
		return 2
	}
	if bytes.Equal(name, []byte("h3")) {
		return 3
	}
	if bytes.Equal(name, []byte("h4")) {
		return 4
	}
	if bytes.Equal(name, []byte("h5")) {
		return 5
	}
	if bytes.Equal(name, []byte("h6")) {
		return 6
	}
	return 0
}
