package author

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthor(t *testing.T) {
	Convey("Format", t, func() {
		a := new(Author)

		So(a.Format(), ShouldEqual, ErrAuthorNoName)

		a = &Author{Name: "fuxiaohei"}
		So(a.Format(), ShouldBeNil)

		So(a.Name, ShouldEqual, a.Nick)
		So(a.Avatar, ShouldBeEmpty)

		a.Email = "fuxiaohei@fuxiaohei.me"
		So(a.Format(), ShouldBeNil)
		So(a.Avatar, ShouldNotBeEmpty)
	})

	Convey("Group", t, func() {
		authors := []*Author{{Name: "abc"}, {Name: "xyz"}, {Name: "123"}}
		group := Group(authors)

		So(group.Get("abc"), ShouldNotBeNil)
		So(group.Get("aaa"), ShouldBeNil)
	})
}
