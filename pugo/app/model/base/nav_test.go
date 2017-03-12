package base

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/go-xiaohei/pugo/app/helper/i18n"
	. "github.com/smartystreets/goconvey/convey"
)

var navI18nString = `
[nav]
home = "Home"
archive = "Archive"
about = "About"
`

var navItemsString = `
[[nav]]
link = "/"
title = "Home"
i18n = "home"
hover = "index"
blank = false

[[nav]]
link = "/archive.html"
title = "Archive"
i18n = "archive"
hover = "archive"

[[nav]]
link = "/about.html"
title = "About-PuGo"
i18n = ""
hover = "about"

[[nav]]
link = "http://pugo.io"
title = "Source"
i18n = "source"
hover = "source"
`

var navWrongString = `
[[nav]]
link = ""
title = ""
i18n = ""
hover = ""
`

func TestNav(t *testing.T) {
	Convey("Nav", t, func() {
		in, err := i18n.New("en", []byte(navI18nString))
		So(err, ShouldBeNil)

		type v struct {
			Group NavGroup `toml:"nav"`
		}
		value := new(v)
		err = toml.Unmarshal([]byte(navItemsString), value)
		So(err, ShouldBeNil)
		So(value.Group, ShouldHaveLength, 4)

		So(value.Group.Format(), ShouldBeNil)

		nav := value.Group[0]
		So(nav.Tr(in), ShouldEqual, "Home")

		nav = value.Group[1]
		So(nav.TrLink(in), ShouldEqual, "en/archive.html")
		So(nav.TrTitle(in), ShouldEqual, "Archive")

		nav = value.Group[2]
		So(nav.TrLink(in), ShouldEqual, "/about.html")
		So(nav.TrTitle(in), ShouldEqual, "About-PuGo")

		nav = value.Group[3]
		So(nav.TrLink(in), ShouldEqual, "http://pugo.io")

		value = new(v)
		err = toml.Unmarshal([]byte(navWrongString), value)
		So(err, ShouldBeNil)
		So(value.Group, ShouldHaveLength, 1)

		So(value.Group.Format(), ShouldEqual, ErrNavInvalid)
	})
}
