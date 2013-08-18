package controllers


import (
	"github.com/astaxie/beego"
)

type Object interface{}

type BaseController struct{
	beego.Controller
}

func (this *BaseController) json(v Object){
    this.Data["json"] = v
    this.ServeJson(true)
}