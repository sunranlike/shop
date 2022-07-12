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
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	//fmt.Printf("%#v", managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

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
	//获取管理员
	id, err := models.ToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据出错", "admin/manager")
		return
	}
	manager := models.Manager{Id: id}

	models.DB.Find(&manager)

	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	//从form表单中获取数据
	id, err1 := models.ToInt(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := models.ToInt(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), "")
	password := strings.Trim(c.PostForm("password"), "")
	email := strings.Trim(c.PostForm("email"), "")
	mobile := strings.Trim(c.PostForm("mobile"), "")
	//执行修改

	manager := models.Manager{Id: id}

	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	//对mobile长度进行判断:
	if len(mobile) > 11 {
		con.Error(c, "电话号码过长", "/admin/manager")
		return
	}
	manager.RoleId = roleId
	//密码填写框是否为空
	if password != "" {
		//长度判断
		if len(password) < 6 {
			con.Error(c, "密码太短了", "/admin/manager/edit?id="+models.ToString(id))
			return
		}
		manager.Password = models.Md5(password) //合规密码设置
	}
	err := models.DB.Save(&manager).Error
	if err != nil {
		con.Error(c, "保存失败", "/admin/manager")
		return
	}
	con.Success(c, "保存数据成功", "/admin/manager")

}

func (con ManagerController) Delete(c *gin.Context) {
	id := c.Query("id") //区分query表单和form表单
	//fmt.Println(id)
	n, err := models.ToInt(id)
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/manager")
		return
	} else {
		//
		manager := &models.Manager{Id: n}
		models.DB.Delete(&manager)

		con.Success(c, "删除成功", "/admin/manager")
	}
}
