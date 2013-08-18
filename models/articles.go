package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

type EnterArticles struct {
    Id int `orm:"auto"json:"id"`
    Title string `orm:"size(256)"json:"title"`
    Uid int `json:"uid"`
    Utype int `orm:"default(0)"json:"user_type"`
    Keywords string `orm:"size(256),null"json:"keywords"`
    Content string `json:"content"`
    Created int `json:"created"`
}


func (this *EnterArticles) Read(id int) error {
    this.Id = id
    err := DefaultOrm.Read(this)
    return err
}

func (this *EnterArticles) Insert() bool {
    retval := false
    this.Created = int(time.Now().Unix())
    DefaultOrm.Insert(this)
    return retval
}

func init() {
    orm.RegisterModel(new(EnterArticles))
}
