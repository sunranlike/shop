package admin

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	//验证md5
	//fmt.Println(models.Md5("123456"))

	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {
	//dologin干嘛了?
	//doLoginin是在点击login按钮的时候的对应操作,点击login操作时,会上传form表单,表单里面有id,password,验证码id（隐藏的）,用户输入的验证码
	//会对验证码就进行验证
	captchaId := c.PostForm("captchaId")
	username := c.PostForm("username") //从form表单中 获取name和密码
	password := c.PostForm("password") //从form表单中获取还没有md5加密的明文密码
	fmt.Println(username, password)
	//从Form表单中寻找captchaID

	verifyValue := c.PostForm("verifyValue") //同上,都是要用来验证的
	fmt.Println("func验证码id:", captchaId, "验证值:", verifyValue)
	//1.验证码判断
	//fmt.Println(verifyValue, captchaId)
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag { //验证,调用验证方法,先对验证码验证再去判断密码

		//2.如果验证成果,还要判断用户是否存在
		userinfo := []models.Manager{}   //操作orm,首先需要实例化你要操作的table的表
		password := models.Md5(password) //md5加密密码,因为底层的密码都是md5的

		models.DB.Where("username=? AND password=?", username, password).Find(&userinfo)

		fmt.Println(userinfo)
		if len(userinfo) > 0 { //如果拿到了数据,执行登录
			//3.执行登录,保存用户信息,执行跳转
			//用cookie或者用session,这里用了session
			session := sessions.Default(c)
			//把结构体转换成json字符串
			userinfoslice, _ := json.Marshal(userinfo)
			session.Options(sessions.Options{ //选项者模式,遍历选项
				MaxAge: 3600,
			})

			session.Set("userinfo", string(userinfoslice)) //设置一个session,是key-value类型,但是set只能保存字符串
			session.Save()
			con.Success(c, "验证码验证成功", "/admin")
		} else {
			con.Error(c, "用户名或者密码错误", "/admin/login")
		}

	} else {
		con.Error(c, "验证码验证失败", "/admin/login")
	}

}
func (con LoginController) Captcha(c *gin.Context) {
	//这个控制器干嘛的?我们的loginin页面会自动访问验证码的url,会调用这个handler
	//这个handler用来生成验证码(makeCaptcha),这个生成的验证码会以json的格式返回id和编码图片
	//于是我们的login页面就可以根据这个编码图片去渲染,当填好验证码并且点击loginin时,会根据id去存储系统判断验证码对不对(后台根据id取得对应的verifyValue)
	id, b64s, err := models.MakeCaptcha() //生产验证码
	//fmt.Println(id, b64s)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{ //以json格式返回验证码,注意要和前端的字段保持一致!!!,因为前端就是根据关键字去获取
		"captchaId":    id,
		"captchaImage": b64s,
	})

}

func (con LoginController) LoginOut(c *gin.Context) {
	//删除session
	session := sessions.Default(c)
	session.Delete("userinfo")
	err := session.Save()
	if err != nil {
		return
	}

	con.Success(c, "退出成功", "/admin/login")
}
