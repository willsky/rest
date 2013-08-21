package controllers

import (
    "rest/models"
    "libs"
    "strconv"
    "fmt"
)

type Article struct {
    BaseController
}

func (this *Article) Post() {
    article := new(models.EnterArticles)
    article.Title = this.GetString("title")
    article.Keywords = this.GetString("keywords")
    article.Content = this.GetString("content")

    if uid, err := this.GetInt("uid"); nil == err && uid > 0 {
        fmt.Println(uid)
        article.Uid = uid
    } else {
        this.jsonResponse(601)
        return
    }

    if utype, err := this.GetInt("user_type"); nil == err {
        article.Utype = uint8(utype)
    }

    if article.Insert() {
        this.jsonResponse(0, article)
    } else {
        this.jsonResponse(505)
    }
}

func (this *Article) Get() {
    article := new(models.EnterArticles)
    var err error
    id, err := strconv.ParseInt(this.Ctx.Params[":objectId"], 10, 64)

    if (nil != err || id < 1) {
        var (
            articles []*models.EnterArticles
            offset int64 = 0
            limit int = 10
        )
        query := make(map[string]libs.Object)

        if v, err := this.GetInt("offset"); nil == err {
           offset = v
        }

        if v, err := this.GetInt("limit"); nil == err {
           limit = int(v)
        }

        query["offset"] = offset
        query["limit"] = limit
        articles, err = article.GetList(query)
        this.jsonResponse(0, articles)
        return
    } else {
        err = article.Read(id)

        if nil != err {
            this.jsonResponse(404)
            return
        }

        this.jsonResponse(0, article)
    }
}


