package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{} //basecontroller 是一个基础的控制器,其他控制器可以类型嵌入

func (con BaseController) Success(c *gin.Context, message string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) Error(c *gin.Context, message string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
