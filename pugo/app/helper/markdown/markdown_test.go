package markdown

import (
	"testing"

	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarkdown(t *testing.T) {
	Convey("Render", t, func() {
		fileData, err := ioutil.ReadFile("markdown_test.md")
		So(err, ShouldBeNil)

		resultData, err := ioutil.ReadFile("markdown_test.html")
		So(err, ShouldBeNil)

		data := Render(fileData)
		So(string(data), ShouldEqual, string(resultData))
	})
}
