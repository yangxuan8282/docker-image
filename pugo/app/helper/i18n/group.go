package i18n

import (
	"errors"
	"sync"
)

var (
	// ErrI18nGroupListMissing means no list data in i18n group
	ErrI18nGroupListMissing = errors.New("i18n group : list is missing")
)

type (
	// Group is a group of i18n objects
	Group struct {
		i18nData map[string]*I18n
		lock     sync.Mutex
		List     []*Item `toml:"lang"`
	}
	// Item is a item description of i18n object in group
	Item struct {
		Lang string `toml:"lang"`
		Name string `toml:"name"`
		File string `toml:"file"`
	}
)

// NewGroup creates a i18n group
func NewGroup() *Group {
	return &Group{
		i18nData: make(map[string]*I18n),
		List:     make([]*Item, 0),
	}
}

// Set sets i18n object to group
func (g *Group) Set(in *I18n) {
	g.lock.Lock()
	g.i18nData[in.Lang] = in
	g.lock.Unlock()
}

// Get gets i18n object by language name
func (g *Group) Get(lang string) *I18n {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.i18nData[lang]
}

// GetItem gets i18n item by language name or i18n object
func (g *Group) GetItem(v interface{}) *Item {
	if str, ok := v.(string); ok {
		for _, item := range g.List {
			if item.Lang == str {
				return item
			}
		}
	}
	if in, ok := v.(*I18n); ok {
		for _, item := range g.List {
			if item.Lang == in.Lang {
				return item
			}
		}
	}
	return new(Item)
}

// Len returns the numbers of i18n in this group
func (g *Group) Len() int {
	return len(g.i18nData)
}

// Has returns whether language's i18n object in this group
func (g *Group) Has(lang string) bool {
	return g.i18nData[lang] != nil
}

// Names returns language names in this group
func (g *Group) Names() []string {
	var s []string
	for key := range g.i18nData {
		s = append(s, key)
	}
	return s
}

// Validate checks the i18n group is corrent data
func (g *Group) Validate() error {
	if len(g.List) == 0 {
		return ErrI18nGroupListMissing
	}
	return nil
}
