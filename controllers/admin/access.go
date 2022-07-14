package admin

import (
	"ginshop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	//fmt.Printf("%#v\n", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})

}
func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块:
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}
func (con AccessController) DoAdd(c *gin.Context) {

	moduleName := strings.Trim(c.PostForm("module_name"), "")
	accessType, err1 := models.Int(c.PostForm("type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入数据出错", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}
	access := models.Access{

		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		con.Error(c, "保存数据失败", "/admin/access/add")
		return
	}

	con.Success(c, "增加数据成功", "/admin/access/add")

	//c.String(http.StatusOK, "doAdd")
}
func (con AccessController) Edit(c *gin.Context) {
	//获取要修改的数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)

	//获取顶级模块:
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})
}
func (con AccessController) DoEdit(c *gin.Context) {
	id, err5 := models.Int(c.PostForm("id"))
	if err5 != nil {
		con.Error(c, "传入key出错", "/admin/access")
		return
	}
	module_name := strings.Trim(c.PostForm("module_name"), "")
	accessType, err1 := models.Int(c.PostForm("type"))
	if err1 != nil {
		con.Error(c, "传入type出错1", "/admin/access")
		return
	}
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	if err2 != nil {
		con.Error(c, "传入module_id出错", "/admin/access")
		return
	}
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")

	if err3 != nil || err4 != nil {
		con.Error(c, "传入数据出错2", "/admin/access")
		return
	}
	if module_name == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+models.String(id))
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = module_name
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.Sort = sort
	access.Description = description
	access.Status = status
	access.ModuleId = moduleId

	err := models.DB.Save(&access).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/access/edit?id="+models.String(id))
		return
	}

	con.Success(c, "修改数据成功", "/admin/access/edit?id="+models.String(id))
	//c.String(http.StatusOK, "doEdit")
}

func (con AccessController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id")) //区分query表单和form表单
	//fmt.Println(id)
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/access")
		return
	} else {
		//
		access := &models.Access{Id: id}

		models.DB.Find(&access)
		if access.ModuleId == 0 { ///判断是否是顶级
			//是顶级模块,还要看有没有子模块
			accessList := []models.Access{}
			models.DB.Where("module_id=?", access.Id).Find(&accessList)
			if len(accessList) > 0 {
				con.Error(c, "当前模块还有子权限,请删除子权限后再删除", "/admin/access")
				return
			}
		} else { //不是顶级模块,是个子模块,直接删除就行
			models.DB.Delete(&access)
		}

		con.Success(c, "删除成功", "/admin/access")
	}
}
