package routers

import (
	"hdy/shiyanshiv/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/data", &controllers.MainController{},"*:GetData")
	beego.Router("/update", &controllers.MainController{},"*:UpdateData")
	beego.Router("/upload", &controllers.UploadController{})
    beego.Router("/", &controllers.MovieController{},"*:GetHtml")
	ns:=beego.NewNamespace("/movie",
		beego.NSRouter("/",&controllers.MovieController{},"*:GetHtml"),
		beego.NSRouter("/cc",&controllers.MovieController{},"*:CC"),
		beego.NSRouter("/sou",&controllers.MovieController{},"*:Search"),
		beego.NSRouter("/video",&controllers.MovieController{},"*:Video"),
		beego.NSRouter("/player",&controllers.MovieController{},"*:Player"))
	beego.AddNamespace(ns)
}
