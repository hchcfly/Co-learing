package controllers


type AboutMeController struct {
	BaseController
}

func (this *AboutMeController) Get() {
	this.Data["wechat"] = "微信:hc1844382037"
	this.Data["qq"] = "QQ:1844382037"
	this.Data["tel"] = "Tel:152954545"
	this.TplName = "aboutme.html"
}