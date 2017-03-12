package base

import (
	"testing"

	"github.com/BurntSushi/toml"
	. "github.com/smartystreets/goconvey/convey"
)

var metaTomlString = `
[meta]
title = "PuGo"
subtitle = "Just for Writing"
keyword = "PuGo,Golang,Static,Website"
desc = "PuGo is a simple static site generator"
root = "http://localhost:9899/"
lang = "en"
`

func TestMeta(t *testing.T) {
	Convey("Meta", t, func() {
		m := new(Meta)
		err := toml.Unmarshal([]byte(metaTomlString), m)
		So(err, ShouldBeNil)

		_, err = m.RootURL()
		So(err, ShouldBeNil)
	})
}
