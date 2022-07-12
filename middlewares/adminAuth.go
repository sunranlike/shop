package middlewares

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//这个中间件会判断访问者的浏览器有没有session，如果没有session或者session不存在，则强制redirect到/admin/login中
	//如果有session，那就什么都不做，
	//当然这个中间件也要对直接访问后台管理系统而没有登录的请求重定向
	fmt.Println("InitAdminAuthMiddleware")
	//进行权限判断 没有登录的用户 不能进入后台管理中心
	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653，但是要从问号处分割，因为
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// 2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo") //调用get方法会获取kv中
	//类型断言 来判断 userinfo是不是一个string,因为get方法返回的是一个空接口
	userinfoStr, ok := userinfo.(string)

	if ok { //如果是字符串：说明有session
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else { //没有session，直接重定向跳转到登陆界面
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
}
