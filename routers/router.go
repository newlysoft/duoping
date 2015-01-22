package routers

import (
	"github.com/astaxie/beego"
	"github.com/newlysoft/duoping/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("admin/login", &controllers.AdminController{}, "*:Login")
}
