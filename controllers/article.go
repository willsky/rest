package controllers

import (
	"rest/models"
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
    this.Data["json"] = make(map[string]int, 1)
    this.ServeJson()
}

func (this *Article) Get() {
    article := new(models.EnterArticles)
    var retval *models.BaseJson = &models.BaseJson{Code: 404, Msg: "Not Found", Data: make(map[string]Object)}
    id, err := this.GetInt("id")
    

    {

        if (nil != err) {
            id = 0
            goto end
        }

        err = article.Read(int(id))
        

        if nil == err {
            retval.Data = article
        }        
    }

    end:
    this.json(retval)
}


