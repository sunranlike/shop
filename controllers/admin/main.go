package admin

import (
	"encoding/json"
	"ginshop/models"
	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	//why 这里被注释掉了?这里的逻辑是重复,这里的主要逻辑就是一个验证session,所以我们把这个逻辑给抽离掉了,也就配置到了group的中间件中
	//这个中间件判断是否有session,如果有才会继续,没有会跳转到login,让你登录.
	session := sessions.Default(c)      //在main页面对session进行获取和判断
	userinfo := session.Get("userinfo") //获取的是json字符串,session的value不能是结构体的切片,因此转换成了json,这里获得的也是json

	userinfoStr, ok := userinfo.(string)

	if ok {
		//step1:获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		//step2:获取用户所有的权限
		accessList := []models.Access{}
		//关联外表
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList) //获取顶级模块
		//查询当前角色对应的权限,并且存为一个id,判断当前id是否在map中,如果在的话就加一个checked属性
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess) //获取当前id对应的AccessId

		roleAccessMap := make(map[int]int) //创建map方式1
		//roleAccessMap := map[int]int{} //创建map方式2
		for _, v := range roleAccess { //把查询到的RoleAccess放入一map中
			roleAccessMap[v.AccessId] = v.AccessId //添加check属性
		}
		//循环遍历accesslist,将当前用户的权限赋值给true
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

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "admin/login")
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

//公告修改状态方法

func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// status = ABS(0-1)   1

	// status = ABS(1-1)  0

	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

//公共修改状态的方法
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}
