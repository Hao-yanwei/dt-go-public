package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.IndexController{})
}
