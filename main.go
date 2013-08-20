package main

import (
	"github.com/astaxie/beego"
	"rest/controllers"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	beego.RESTRouter("/article", &controllers.Article{})
	beego.Run()
}

func init() {
    orm.RunCommand()
}
