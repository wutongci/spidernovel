package components

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	//读取不同环境数据库配置
	env := beego.BConfig.RunMode
	if env == "dev" {
		mysqluser := beego.AppConfig.String(env + "::mysqluser")
		mysqlpass := beego.AppConfig.String(env + "::mysqlpass")
		mysqlurls := beego.AppConfig.String(env + "::mysqlurls")
		mysqlport := beego.AppConfig.String(env + "::mysqlport")
		mysqldb := beego.AppConfig.String(env + "::mysqldb")
		//初始化数据库
		orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8mb4")
	}
	//设置最大链接数
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 300)
}
