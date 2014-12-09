package routers

import (
	"github.com/newlysoft/duoping/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
