package demo1

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
		// So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return true when a == nil  && b == nil", t, func() {
		So(StringSliceEqual(nil, nil), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return false when a == nil  && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when a != nil  && b != nil", t, func() {
		a := []string{"hello", "world"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}

// Convey语句可以无限嵌套 以体现测试用例之间的关系
// 需要注意的是 只有最外层的Convey需要传入*testing.T 类型的变量t
func TestStringSliceEqual1(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		Convey("should return true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}

func TestSummer(t *testing.T) {
	Convey("TestSummer", t, func() {
		So("summer", ShouldSummerBeComming, "comming")
		So("winter", ShouldSummerBeComming, "comming")
	})
}
