package routers

import (
	"study/day20_beego/WEB/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/test_request", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test_Session", &controllers.TestSession{}, "get:Get;post:Post")
	beego.Router("/test_model", &controllers.TestModelController{}, "get:Get;post:Post")
}
