package i18n

import (
	"fmt"
	"strings"

	"io/ioutil"
	"path"

	"github.com/BurntSushi/toml"
)

type (
	I18n struct {
		values map[string]map[string]string
		Lang   string
	}
)

// Tr converts string
func (i *I18n) Tr(str string) string {
	strSlice := strings.Split(str, ".")
	if len(strSlice) != 2 {
		return str
	}
	if m, ok := i.values[strSlice[0]]; ok {
		if v := m[strSlice[1]]; v != "" {
			return v
		}
	}
	return str
}

// Trf converts string with arguments
func (i *I18n) Trf(str string, values ...interface{}) string {
	return fmt.Sprintf(i.Tr(str), values...)
}

// NewEmpty creates new empty i18n object,
// it will keep i18 tool working, but no translated value
func NewEmpty() *I18n {
	return &I18n{
		Lang:   "nil",
		values: make(map[string]map[string]string),
	}
}

// New creates i18n object with data and language name
func New(lang string, data []byte) (*I18n, error) {
	maps := make(map[string]map[string]string)
	if err := toml.Unmarshal(data, &maps); err != nil {
		return nil, err
	}
	for k, v := range maps {
		if len(v) == 0 {
			return nil, fmt.Errorf("i18n section '%s' is empty", k)
		}
	}
	return &I18n{Lang: lang, values: maps}, nil
}

// NewFromFile creates i18n object from file,
// filename is language name
func NewFromFile(file string) (*I18n, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	name := path.Base(file)
	name = strings.TrimSuffix(name, path.Ext(name))
	return New(name, data)
}

// LangCode returns correct language code possibly
// en-US -> [en-US,en-us,en]
func LangCode(lang string) []string {
	languages := []string{lang} // [en-US]
	lower := strings.ToLower(lang)
	if lower != lang {
		languages = append(languages, lower) // use lowercase language code, [en-us]
	}
	if strings.Contains(lang, "-") {
		languages = append(languages, strings.Split(lang, "-")[0]) // use first word if en-US, [en]
	}
	return languages
}
