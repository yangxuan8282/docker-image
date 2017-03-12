package i18n

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	i18nCorrectData = `# navigation
[nav]
home = "Home"
archive = "Archive"
about = "About"

# about post
[post]
list = "All Posts"
archive = "Archives"
tag = "Tags"
readmore = "Read More"

# about pagination
[pager]
prev = "Prev"
next = "Next"
current = "Current-%d"

[page]
index = "Index"`
	i18nIncorrectData = `[nav]]]]]
home = "Home"
archive = "Archive"
about = "About"
`
	i18nEmptySectionData = `# navigation
[nav]
home = "Home"
archive = "Archive"
about = "About"

# about post
[post]
`
)

func TestEmpty(t *testing.T) {
	Convey("Empty", t, func() {
		empty := NewEmpty()
		So(empty.Lang, ShouldEqual, "nil")
		So(empty.values, ShouldHaveLength, 0)
	})
}

func TestNew(t *testing.T) {
	Convey("New", t, func() {
		in, err := New("en", []byte(i18nCorrectData))
		So(err, ShouldBeNil)
		So(in.Lang, ShouldEqual, "en")
		So(in.values, ShouldHaveLength, 4) // four section

		Convey("Tr & Trf", func() {
			So(in.Tr("nav.home"), ShouldEqual, "Home")
			So(in.Trf("pager.current", 10), ShouldEqual, "Current-10")
			So(in.Tr("page"), ShouldEqual, "page")
			So(in.Tr("nav.invalid"), ShouldEqual, "nav.invalid")
		})

		Convey("NewError", func() {
			_, err := New("en", []byte(i18nIncorrectData))
			So(err, ShouldNotBeNil)

			_, err = New("en", []byte(i18nEmptySectionData))
			So(err, ShouldNotBeNil)
		})
	})
}

func TestNewFile(t *testing.T) {
	Convey("NewFile", t, func() {
		in, err := NewFromFile("i18n_test.toml")
		So(err, ShouldBeNil)
		So(in.Lang, ShouldEqual, "i18n_test")

		_, err = NewFromFile("missing.toml")
		So(err, ShouldNotBeNil)
	})
}

func TestLangCode(t *testing.T) {
	Convey("LangCode", t, func() {
		codes := LangCode("en-US")
		So(codes, ShouldHaveLength, 3)
		So(codes, ShouldContain, "en")
		So(codes, ShouldContain, "en-us")
		So(codes, ShouldContain, "en-US")

		codes = LangCode("zh")
		So(codes, ShouldHaveLength, 1)
	})
}

func TestGroup(t *testing.T) {
	Convey("Group", t, func() {
		g := NewGroup()
		in, err := New("en", []byte(i18nCorrectData))
		So(err, ShouldBeNil)

		in2, err := NewFromFile("i18n_test.toml")
		So(err, ShouldBeNil)

		g.Set(in)
		g.Set(in2)

		So(g.Len(), ShouldEqual, 2)
		So(g.Names(), ShouldContain, "i18n_test")
		So(g.Get("en"), ShouldNotBeNil)
		So(g.Get("zh"), ShouldBeNil)
		So(g.Has("jp"), ShouldEqual, false)
	})
}
