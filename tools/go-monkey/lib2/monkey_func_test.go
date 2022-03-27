package lib2

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/fake"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	outputExpect = "xxx-vethName100-yyy"
)

func TestApplyFunc(t *testing.T) {
	Convey("TestApplyFunc", t, func() {

		Convey("one func for success", func() {
			patches := ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return outputExpect, nil
			})
			defer patches.Reset()
			output, err := fake.Exec("", "")
			So(err, ShouldEqual, nil)
			So(output, ShouldEqual, outputExpect)
		})

	})
}
