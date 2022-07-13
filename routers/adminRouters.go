package routers

import (
	"ginshop/controllers/admin"
	"ginshop/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	//这里配置的是中间件,相当于这个url下的所有访问，都会走一下这个中间件,这也是一个逻辑业务的抽离,因为整个admin其实都需要验证身份(当然个别url除外：验证码或者本身就在login页面)
	//所以我们配置一个全局中间件,这个中间件就是对用户的session去判断,是否有session已登录,如果没有session,那就要跳转到login界面
	//这种全局中间件也是一种抽象的概念,将公共的部分抽离出来,避免重复劳动.
	{
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		adminRouters.POST("/role/doadd", admin.RoleController{}.DoAdd)

		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/role/doedit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)
		//this is a test

		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)

		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)
	}
}
