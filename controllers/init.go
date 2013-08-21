package controllers

import (
    "github.com/astaxie/beego"
    "rest/models"
    "rest/libs"
)

type BaseController struct{
    beego.Controller
}

func (this *BaseController) json(v libs.Object){
    this.Data["json"] = v
    this.ServeJson(true)
}

func (this *BaseController) jsonResponse(status int, v ... libs.Object) {
    var (
        data *models.BaseJson = new(models.BaseJson)
        message string
    )

    data.Code = status

    if len(v) > 0 {
        data.Data = v[0]
    } else {
        data.Data = make(map[string]int)
    }


    switch status {
    case 0:
        message = "OK"
    case 404:
        message = "Not fount"
    case 501:
        message = "Server internal exception"
    case 601:
        message = "Request error"
    }

    data.Msg = message
    this.json(data)
}

