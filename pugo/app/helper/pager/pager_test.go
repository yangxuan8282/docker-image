package pager

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPager(t *testing.T) {
	Convey("Cursor", t, func() {
		cursor := NewCursor(10, 100)
		So(cursor.pages, ShouldEqual, 10)
		So(cursor.size, ShouldEqual, 10)
		So(cursor.all, ShouldEqual, 100)

		cursor = NewCursor(10, 101)
		So(cursor.pages, ShouldEqual, 11)

		Convey("Pager", func() {
			pager := cursor.Page(2)
			So(pager.Current, ShouldEqual, 2)
			So(pager.Begin, ShouldEqual, 10)
			So(pager.End, ShouldEqual, 20)
			pager.SetLayout("page-%d")
			So(pager.PrevURL(), ShouldEqual, "page-1")
			So(pager.URL(), ShouldEqual, "page-2")
			So(pager.NextURL(), ShouldEqual, "page-3")

			for i, item := range pager.Items() {
				So(item.Page, ShouldEqual, i+1)
				So(item.Link, ShouldEqual, fmt.Sprintf("page-%d", i+1))
			}

			pager = cursor.Page(11)
			So(pager.Next, ShouldEqual, 0)
			So(pager.Prev, ShouldEqual, 10)
			So(pager.NextURL(), ShouldEqual, "")

			pager = cursor.Page(1)
			So(pager.PrevURL(), ShouldEqual, "")

			pager = cursor.Page(0)
			So(pager, ShouldBeNil)
			pager = cursor.Page(99)
			So(pager, ShouldBeNil)
		})
	})
}
