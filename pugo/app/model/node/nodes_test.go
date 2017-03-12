package node

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNode(t *testing.T) {
	Convey("Node", t, func() {
		tree := NewTree("")
		tree.Add("", "Nil", NodePage, 1)
		tree.Add("about/me", "Me", NodePage, 1)
		tree.Add("about", "About", NodePage, 3)
		tree.Add("/archive", "Archive", NodePost, 2)
		tree.Add("/user", "User", NodeXML, 1)
		tree.Add("user/fuxiaohei", "Fuxiaohei", NodePage, 1)
		tree.SortChildren()
		tree.Print("")

		So(tree.Children, ShouldHaveLength, 3)
		So(tree.Children[0].Title, ShouldEqual, "User")
		So(tree.Children[2].Title, ShouldEqual, "About")
		So(tree.Len(), ShouldEqual, 5)
	})
}
