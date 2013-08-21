package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "libs"
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
    _, err := DefaultOrm.Insert(this)

    if nil == err {
        retval = true
    }

    return retval
}

func (this *EnterArticles) GetList(query map[string]libs.Object) (articles []*EnterArticles, err error) {
    var (
        offset int64 = 0
        limit int = 10
    )

    if val, ok := query["offset"]; ok {
        if val, t := val.(int64); t {
            offset = val
        }

        if val, t := val.(int); t {
            offset = int64(val)
        }

        if offset < 0 {
            offset = 0
        }
    }

    if val, ok := query["limit"]; ok {
        if val, t := val.(int); t {
            limit = val
        }

        if limit < 1 {
            limit = 10
        }
    }
    _, err = DefaultOrm.QueryTable(this).OrderBy("-id").Offset(offset).Limit(limit).All(&articles)
    return
}

func init() {
    orm.RegisterModel(new(EnterArticles))
}
