package admin

import (
	"fmt"
	"ginshop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	fmt.Println(" ")
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), "")
	description := strings.Trim(c.PostForm("description"), "")

	if title == "" || description == "" {
		con.Error(c, "不能为空", "admin/role/add") //调用继承的方法
	} //如果不为空,要加role了
	role := models.Role{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "创建失败", "/admin/role/add")
		return
	} else {
		con.Success(c, "增加成功", "/admin/role")
	}

	//c.String(http.StatusOK, "doAdd")
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.ToInt(c.Query("id"))

	if err != nil {
		con.Error(c, "传入数据出错", "admin/role")
		return
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
	//c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{})
}
func (con RoleController) DoEdit(c *gin.Context) {
	id, err1 := models.ToInt(c.PostForm("id"))
	//fmt.Println(id)
	if err1 != nil {
		con.Error(c, "数据错误", "admin/role")
		return
	}

	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "角色titile不能为空", "admin/role/edit")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	//查询数据
	role.Title = title
	role.Description = description

	err2 := models.DB.Save(&role).Error //修改数据记得save
	if err2 != nil {
		con.Error(c, "gorm失败", "/admin/role/edit?id="+models.ToString(id)) //重新跳转
		return
	} else {
		con.Success(c, "修改成功", "/admin/role/edit?id="+models.ToString(id))
		//踩过的坑:注意一定要最前面加一个"/",否则就是当前域加上这个url,例如admin/role/admin/role/edit?id=,明显有悖于结论
	}

	//c.String(http.StatusOK, "doEdit")
}

func (con RoleController) Delete(c *gin.Context) {
	id := c.Query("id") //区分query表单和form表单
	//fmt.Println(id)
	n, err := models.ToInt(id)
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/role")
		return
	} else {
		//
		role := &models.Role{Id: n}
		models.DB.Delete(&role)

		con.Success(c, "删除成功", "/admin/role")
	}
}

func (con RoleController) Auth(c *gin.Context) {
	id, err := models.ToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"id": id,
	})
	//c.String(http.StatusOK, "Auth", gin.H{})
}
func (con RoleController) DoAuth(c *gin.Context) {
	//id, err := models.ToInt(c.Query("id"))
	c.String(http.StatusOK, "DoAuth")
}
