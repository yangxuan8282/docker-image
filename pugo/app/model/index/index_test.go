package index

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIndex(t *testing.T) {
	Convey("Index", t, func() {
		data, _ := ioutil.ReadFile("index_test.html")
		idx := New(data)
		So(idx, ShouldHaveLength, 1)

		So(idx[0].Title, ShouldEqual, "Go 1.8 Release Notes")
		So(idx[0].Children, ShouldHaveLength, 7)

		So(idx[0].Children[0].Children, ShouldHaveLength, 0)
		So(idx[0].Children[1].Children, ShouldHaveLength, 0)
		So(idx[0].Children[2].Children, ShouldHaveLength, 1)
		So(idx[0].Children[3].Children, ShouldHaveLength, 11)
		So(idx[0].Children[4].Children, ShouldHaveLength, 3)
		So(idx[0].Children[5].Children, ShouldHaveLength, 3)
		So(idx[0].Children[6].Children, ShouldHaveLength, 5)

		So(idx[0].Children[6].Children[0].Children, ShouldHaveLength, 0)
		So(idx[0].Children[6].Children[1].Children, ShouldHaveLength, 0)
		So(idx[0].Children[6].Children[2].Children, ShouldHaveLength, 1)
		So(idx[0].Children[6].Children[3].Children, ShouldHaveLength, 0)
		So(idx[0].Children[6].Children[4].Children, ShouldHaveLength, 0)

		So(idx[0].Children[1].Anchor, ShouldEqual, "language")
	})
}
