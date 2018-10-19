package main

import (
	_ "spidernovel/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
)

func main() {
	//logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"/home/wwwlogs/book/book.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":1000}`)
	////异步
	//logs.Async()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
