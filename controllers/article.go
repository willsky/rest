package controllers

import (
	"rest/models"
    "fmt"
)

type Article struct {
    BaseController
}

func (this *Article) Post() {
    article := new(models.EnterArticles)
    article.Title = "Test"
    article.Uid = 1
    article.Keywords = "test"
    article.Content = "test test"
    article.Insert()
    this.jsonResponse(0, make(map[string]int))
}

func (this *Article) Get() {
    article := new(models.EnterArticles)
    id, err := this.GetInt("id")

    if (nil != err) {
        id = 0
        this.jsonResponse(601)
        return
    }

    err = article.Read(int(id))

    if nil != err {
        this.jsonResponse(404)
        return
    }

    this.jsonResponse(0, article)
}


