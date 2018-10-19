package controllers

import(
	"github.com/astaxie/beego"
)

type BaseController struct{
	beego.Controller
}

// 固定返回的json数据格式
// code: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) ToJson (code int, msg string, data interface{}){
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
}