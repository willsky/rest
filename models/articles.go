package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

type EnterArticles struct {
    Id int64 `orm:"auto" json:"id"`
    Title string `orm:"size(256)" json:"title"`
    Uid int64 `orm:"index" json:"uid"`
    Utype uint8 `orm:"default(0)" json:"user_type"`
    Keywords string `orm:"size(256);null" json:"keywords"`
    Content string `orm:"type(text)" json:"content"`
    Created uint `json:"created"`
}

func (this *EnterArticles) TableName() string {
    return "articles"
}

func (this *EnterArticles) Read(id int64) error {
    this.Id = id
    err := DefaultOrm.Read(this)
    return err
}

func (this *EnterArticles) Insert() bool {
    retval := false
    this.Created = uint(time.Now().Unix())
    DefaultOrm.Insert(this)
    return retval
}

func init() {
    orm.RegisterModel(new(EnterArticles))
}
