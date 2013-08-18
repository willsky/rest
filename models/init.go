package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

var (
    DefaultOrm orm.Ormer
)

func init() {
    orm.Debug = true
    orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/enterprise_site?charset=utf8", 10)
    DefaultOrm = orm.NewOrm()
}


type BaseJson struct {
	Code int `json:"errcode"`
	Msg string `json:"msg"`
	Data Model `json:"data"`
}

type Model interface {}