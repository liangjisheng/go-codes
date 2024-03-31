package main

import (
	"io/ioutil"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/GoAdminGroup/go-admin/context"

	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:8080/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// customize a plugin
	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	// you can also add config like:
	//
	// import "github.com/GoAdminGroup/go-admin/modules/config""
	//
	// cfg := config.Config{
	//	 Databases: config.DatabaseList{
	//		"default": {
	//			Host:       "127.0.0.1",
	//			Port:       "3306",
	//			User:       "root",
	//			Pwd:        "root",
	//			Name:       "godmin",
	//			MaxIdleCon: 50,
	//			MaxOpenCon: 150,
	//			Driver:     db.DriverMysql,
	//		},
	//	},
	//	UrlPrefix: "admin",
	//	IndexUrl:  "/",
	//	Debug:     true,
	//	Language:  language.CN,
	// }
	//
	// eng.AddConfig(cfg)

	if err := eng.AddConfigFromJSON("./config.json").
		AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// customize your index page
	r.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			contextRes := context.NewContext(nil)
			return datamodel.GetContent(contextRes)
		})
	})

	_ = r.Run("127.0.0.1:8080")
}
