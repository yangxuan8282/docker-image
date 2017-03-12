package theme

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/go-xiaohei/pugo/app/helper/printer"
)

// Meta is description of theme
type Meta struct {
	Name string   `toml:"name"`
	Repo string   `toml:"repo"`
	URL  string   `toml:"url"`
	Date string   `toml:"date"`
	Desc string   `toml:"desc"`
	Tags []string `toml:"tags"`

	MinVersion string `toml:"min_version"`

	// Authors []*model.Author `toml:"author" ini:"-"`
	Refs []*Reference `toml:"ref"`

	License    string `toml:"license" `
	LicenseURL string `toml:"license_url"`
}

// Reference is reference of user links
type Reference struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
	Repo string `toml:"repo"`
}

// NewMeta parses bytes to theme meta
func NewMeta(data []byte) (*Meta, error) {
	meta := new(Meta)
	if err := toml.Unmarshal(data, meta); err != nil {
		return nil, err
	}
	if meta.Name == "" {
		return nil, nil
	}
	return meta, nil
}

// NewMetaFromFile parses theme meta via file
func NewMetaFromFile(file string) (*Meta, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	printer.Logf("read %s", file)
	return NewMeta(data)
}
