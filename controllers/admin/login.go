package admin

import (
	"fmt"
	"ginshop02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {
	captchaId := c.PostForm("captchaId")

	verifyValue := c.PostForm("verifyValue")

	fmt.Println(verifyValue)
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag == true {
		c.String(http.StatusOK, "验证成功", gin.H{
			"captchaId":   captchaId,
			"verifyValue": verifyValue,
		})
	} else {
		c.String(http.StatusOK, "验证失败")
	}

}
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"CaptchaId":     id,
		"CaptchaImages": b64s,
	})

}
