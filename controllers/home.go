package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
