package author

import (
	"errors"

	"github.com/go-xiaohei/pugo/app/helper/gravatar"
)

type (
	// Author is author item in meta file
	Author struct {
		Name    string `toml:"name"`
		Nick    string `toml:"nick"`
		Email   string `toml:"email"`
		URL     string `toml:"url"`
		Avatar  string `toml:"avatar"`
		Bio     string `toml:"bio"`
		Repo    string `toml:"repo"` // github repository
		IsOwner bool   // must be the first author
	}
	// Group is a collection of authors
	Group []*Author
)

var (
	// ErrAuthorNoName means author without name
	ErrAuthorNoName = errors.New("author must have name")
)

// Format formats author's data
func (a *Author) Format() error {
	if a.Name == "" {
		return ErrAuthorNoName
	}
	if a.Nick == "" {
		a.Nick = a.Name
	}
	if a.Avatar == "" && a.Email != "" {
		a.Avatar = gravatar.FromEmail(a.Email, 0)
	}
	return nil
}

// Get gets an author in this group by name
func (g Group) Get(name string) *Author {
	for _, a := range g {
		if a.Name == name {
			return a
		}
	}
	return nil
}
