package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

var (
    DefaultOrm orm.Ormer
)

func init() {

    if "dev" == beego.RunMode {
        orm.Debug = true 
    }

    dns := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
    beego.AppConfig.String("defaultmysqluser"),
    beego.AppConfig.String("defaultmysqlpasswd"),
    beego.AppConfig.String("dbprotocol"),
    beego.AppConfig.String("defaultmysqlhost"),
    beego.AppConfig.String("defaultmysqlport"),
    beego.AppConfig.String("defaultmysqldatbase"),
    beego.AppConfig.String("dbcharset"))

    defaultMaxIdle, _ := beego.AppConfig.Int("defaultmysqlmaxidle")
    orm.RegisterDataBase("default", "mysql", dns, defaultMaxIdle)
    DefaultOrm = orm.NewOrm()
}


type BaseJson struct {
	Code int `json:"errcode"`
	Msg string `json:"msg"`
	Data Model `json:"data"`
}

type Model interface {}
