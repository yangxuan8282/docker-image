package watcher

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWatcher(t *testing.T) {
	counter := 0

	Convey("Watcher", t, func() {
		os.Mkdir("test", os.ModePerm)

		w, err := New()
		So(err, ShouldBeNil)

		err = w.Add("./")
		So(err, ShouldBeNil)

		err = w.Add("xyz")
		So(err, ShouldNotBeNil)

		w.SetExt(".txt")

		go func() {
			time.Sleep(2 * time.Second)
			ioutil.WriteFile("test/abc.txt", []byte("111222333"), os.ModePerm)

			time.Sleep(2 * time.Second)
			ioutil.WriteFile("test/abc.txt", []byte("abcdefg"), os.ModePerm)

			time.Sleep(2 * time.Second)
			os.Remove("test/abc.txt")

			time.Sleep(2 * time.Second)
			w.Close()
		}()

		err = w.Start(nil)
		So(err, ShouldEqual, ErrWatchNoFunction)

		err = w.Start(func() {
			counter++
		})
		So(err, ShouldBeNil)
		So(counter, ShouldEqual, 3)
		os.RemoveAll("test")
	})
}
