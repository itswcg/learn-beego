package routers

import (
	"github.com/astaxie/beego"
	"learn-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
