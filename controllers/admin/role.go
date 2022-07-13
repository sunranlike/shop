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
	//获取角色id
	id, err := models.ToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取所有的权限，
	accessList := []models.Access{}
	//关联外表
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList) //获取顶级模块

	//查询当前角色对应的权限,并且存为一个id,判断当前id是否在map中,如果在的话就加一个checked属性
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", id).Find(&roleAccess) //获取当前id对应的AccessId

	roleAccessMap := make(map[int]int) //创建map方式1
	//roleAccessMap := map[int]int{} //创建map方式2
	for _, v := range roleAccess { //把查询到的RoleAccess放入一map中
		roleAccessMap[v.AccessId] = v.AccessId //添加check属性
	}
	//判判断用户的权限是否在当前用户的列表里面,如果有权限的话,给check设置为true
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     id,
		"accessList": accessList,
	})
	//c.String(http.StatusOK, "Auth", gin.H{})
}
func (con RoleController) DoAuth(c *gin.Context) {
	//获取角色id和权限id
	roleId, err := models.ToInt(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取权限id,是个切片
	accessIds := c.PostFormArray("access_node[]")

	//删除当前角色对应的所有权限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess) //删除数据

	//增加所有选中的权限

	for _, v := range accessIds {
		//增加数据
		roleAccess.RoleId = roleId //form表单中的值赋值给结构体,因为id是唯一的,不是从遍历中获得
		roleAccess.AccessId, _ = models.ToInt(v)
		models.DB.Create(&roleAccess)

	}
	//id, err := models.ToInt(c.Query("id"))
	con.Success(c, "添加权限成功", "/admin/role")
}
