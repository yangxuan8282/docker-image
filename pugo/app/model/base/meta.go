package base

import "net/url"

type (
	// Meta is the meta info of site
	Meta struct {
		Title       string `toml:"title"`
		Subtitle    string `toml:"subtitle"`
		Keyword     string `toml:"keyword"`
		Description string `toml:"desc"`
		Root        string `toml:"root"`
		Lang        string `toml:"lang"`
		SrcFile     string `toml:"-"`
	}
)

// RootURL parses root to url object
func (m *Meta) RootURL() (*url.URL, error) {
	return url.Parse(m.Root)
}
