package gravatar

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFromEmail(t *testing.T) {
	Convey("FromEmail", t, func() {
		avatar := FromEmail("fuxiaohei@vip.qq.com", 60)
		So(avatar, ShouldEqual, "https://www.gravatar.com/avatar/f72f7454ce9d710baa506394f68f4132?size=60")

		avatar = FromEmail("fuxiaohei@vip.qq.com", 0)
		So(avatar, ShouldEqual, "https://www.gravatar.com/avatar/f72f7454ce9d710baa506394f68f4132?size=80")
	})
}
