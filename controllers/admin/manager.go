package admin

import (
	"ginshop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})

}
func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}
func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err := models.ToInt(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "转换角色失败", "/admin/manger/add")
		return
	} else {
		username := strings.Trim(c.PostForm("username"), "")
		password := strings.Trim(c.PostForm("password"), "")
		email := strings.Trim(c.PostForm("email"), "")
		mobile := strings.Trim(c.PostForm("mobile"), "")
		//用户名或者密码长度是否合法
		if len(username) < 2 || len(password) < 6 {
			con.Error(c, "用户名或者密码不合法", "/admin/manger/add")
			return
		}
		//判断管理员是否存在
		mangerList := []models.Manager{}
		models.DB.Where("username=?", username).Find(&mangerList)
		if len(mangerList) > 0 {
			con.Error(c, "管理员已经存在", "/admin/manager/add")
			return
		}
		//如果没有这个管理员,则正常添加成员

		manager := models.Manager{
			//Id:       roleId,
			Username: username,
			Password: models.Md5(password),
			Mobile:   mobile,
			Email:    email,
			Status:   1,
			RoleId:   roleId,
			AddTime:  int(models.GetUnix()),
			IsSuper:  0,
		}
		err = models.DB.Create(&manager).Error
		if err != nil {
			con.Error(c, "存入数据库失败", "/admin/manger/add")
		}
	}
	con.Success(c, "增加成功", "/admin/manager/add")
}

func (con ManagerController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}
func (con ManagerController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
