package main

import (
	"ginshop/models"
	"ginshop/routers"
	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥,存储在redis中
	//注意了,redis我们在验证码和session中都用到了,两个最好要统一
	store, _ := redis.NewStore(10, "tcp", models.RedisAddr+":"+models.RedisProt, models.RedisPassword, []byte("secret"))

	//配置session的全局中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎,这里我们用的是redis存储
	//mysession是名字
	r.Use(sessions.Sessions("mysession", store)) //使用session中间件

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
