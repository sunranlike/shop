package models

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

//时间戳转换成日期,这个方法会mapFunc中使用,可以让前端代码调用
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}
func GetUnixN() int64 {
	return time.Now().UnixNano()
}

//获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))

}
func Sha256(str string) string {
	h := sha256.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func Int(str string) (n int, err error) {
	n, err = strconv.Atoi(str)
	return
}
func String(n int) string {

	str := strconv.Itoa(n)
	return str
}

//上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件,也是从form表单中获取的,
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}
	//限制后缀
	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok { //通过map获取是否合法
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixN(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst) //返回的dst是一个文件的目录
	return dst, nil

}

//表示把string转换成Float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

//把字符串解析成html，这里在前端静态页面有用到，用来吧字符串给转换成html格式的字段，直接插入到html中
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}
