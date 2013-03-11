package main

import (
	"github.com/astaxie/beego"
	"weixin/controllers"
)

func main() {
	beego.Info("come in")
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/weixin", &controllers.MainController{})
	beego.Run()
}
