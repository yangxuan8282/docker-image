package ziper

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	gzipString = "The compiler back end introduced in Go 1.7 for 64-bit x86 is now used on all architectures, and those architectures should see significant performance improvements. For instance, the CPU time required by our benchmark programs was reduced by 20-30% on 32-bit ARM systems. There are also some modest performance improvements in this release for 64-bit x86 systems. The compiler and linker have been made faster. Compile times should be improved by about 15% over Go 1.7. There is still more work to be done in this area: expect faster compilation speeds in future releases."
)

func TestGzip(t *testing.T) {
	Convey("Gzip & UnGzip", t, func() {
		data, err := Gzip([]byte(gzipString))
		So(err, ShouldBeNil)

		srcData, err := UnGzip(data)
		So(err, ShouldBeNil)
		So(string(srcData), ShouldEqual, gzipString)
	})

	Convey("GzipFile & UnGzipFile", t, func() {
		_, err := GzipFile("xyz.txt")
		So(err, ShouldNotBeNil)

		data, err := GzipFile("gzip_test.txt")
		So(err, ShouldBeNil)

		srcData, err := UnGzip(data)
		So(err, ShouldBeNil)
		So(string(srcData), ShouldEqual, gzipString)
	})

	Convey("GzipFileBase64 & UnGzipFileBase64", t, func() {
		_, err := GzipFileBase64("xyz.txt")
		So(err, ShouldNotBeNil)

		data, err := GzipFileBase64("gzip_test.txt")
		So(err, ShouldBeNil)

		err = UnGzipFileBase64(data, "gzip_test.txt2")
		So(err, ShouldBeNil)

		srcData, err := ioutil.ReadFile("gzip_test.txt2")
		So(err, ShouldBeNil)
		So(string(srcData), ShouldEqual, gzipString)

		os.Remove("gzip_test.txt2")
	})
}
