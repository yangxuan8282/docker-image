package page

import (
	"html/template"
	"testing"

	"os"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPage(t *testing.T) {
	Convey("Normal", t, func() {
		p, err := NewFromFile("testdata/page.md", "testdata/page")
		So(err, ShouldBeNil)
		So(p.URL(), ShouldEqual, "testdata/page.html")
		So(p.DstFile(), ShouldEqual, "testdata/page.html")
		So(p.SrcFile(), ShouldEqual, "testdata/page.md")
		So(p.Content(), ShouldHaveLength, 2243)
		So(p.ContentHTML(), ShouldHaveSameTypeAs, template.HTML("abc"))
		So(p.ContentLength(), ShouldEqual, 2243)
		So(p.CreateTime().Unix(), ShouldEqual, 1483272020)
		So(p.UpdateTime().Unix(), ShouldEqual, 1483303800)
		So(p.Author(), ShouldBeNil)
		So(p.Index(), ShouldHaveLength, 2)
	})

	Convey("CreateTime", t, func() {
		info, _ := os.Stat("testdata/no_createtime.md")
		t := info.ModTime()

		p, err := NewFromFile("testdata/no_createtime.md", "testdata/no_createtime")
		So(err, ShouldBeNil)
		So(p.CreateTime().Unix(), ShouldEqual, t.Unix())
		So(p.CreateTime().Unix(), ShouldEqual, p.UpdateTime().Unix())
	})

	Convey("MetaFormat", t, func() {
		_, err := NewFromFile("testdata/unknown_meta.md", "testdata/unknown_meta")
		So(err, ShouldEqual, ErrPageFrontMetaTypeUnknown)
	})

	Convey("TimeFormat", t, func() {
		p := &Page{
			Created: "01-01 13:30",
		}
		So(p.formatTime(), ShouldNotBeNil)
	})
}
