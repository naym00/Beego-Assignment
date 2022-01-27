package routers

import (
	"beego_project/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.CatController{}, "GET:Cat_data")
	beego.Router("/internalprocess", &controllers.CatController{}, "GET:FetchFromAPI")
}