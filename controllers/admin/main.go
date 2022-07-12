package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	//why 这里被注释掉了?这里的逻辑是重复,这里的主要逻辑就是一个验证session,所以我们把这个逻辑给抽离掉了,也就配置到了group的中间件中
	//这个中间件判断是否有session,如果有才会继续,没有会跳转到login,让你登录.
	//session := sessions.Default(c)      //在main页面对session进行获取和判断
	//userinfo := session.Get("userinfo") //获取的是json字符串,session的value不能是结构体的切片,因此转换成了json,这里获得的也是json
	//
	//userinfoStr, ok := userinfo.(string)
	//if ok {
	//	var u []models.Manager
	//	json.Unmarshal([]byte(userinfoStr), &u)
	//	fmt.Println(u)
	//	c.JSON(http.StatusOK, gin.H{
	//		"username": u[0].Username,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"username": "session不存在",
	//	})
	//}

	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
