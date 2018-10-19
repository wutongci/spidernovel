package components

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
)

var Cache  cache.Cache
func init()  {
	env := beego.BConfig.RunMode
	redisconn := beego.AppConfig.String(env + "::redis_conn")
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%d","password":"%s"}`,
		"spiderbook",
		redisconn,
		0,
		"",
	)
	Cache, _ = cache.NewCache("redis", redisConf)
}
