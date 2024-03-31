package middleware

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {
	InitCasbin("rbac_models.conf")
}

func TestCasbin(t *testing.T) {
	var (
		sub, obj, act string
	)

	convey.Convey("test auth", t, func() {

		convey.Convey("visitor", func() {
			convey.Convey("visitor GET", func() {
				sub, obj, act = UserRoleVisitor, "/api/nft/activity", "GET"
				res, _ := enforcer.Enforce(sub, obj, act)
				convey.So(res, convey.ShouldEqual, true)
			})
			convey.Convey("visitor POST", func() {
				sub, obj, act = UserRoleVisitor, "/api/nft/activity", "POST"
				res, _ := enforcer.Enforce(sub, obj, act)
				convey.So(res, convey.ShouldEqual, false)
			})
		})

		convey.Convey("admin", func() {
			convey.Convey("admin GET", func() {
				sub, obj, act = UserRoleAdmin, "/api/nft/activity", "GET"
				res, _ := enforcer.Enforce(sub, obj, act)
				convey.So(res, convey.ShouldEqual, true)
			})
			convey.Convey("admin POST", func() {
				sub, obj, act = UserRoleAdmin, "/api/nft/activity", "POST"
				res, _ := enforcer.Enforce(sub, obj, act)
				convey.So(res, convey.ShouldEqual, true)
			})
		})

	})
}
