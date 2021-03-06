package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

//var store = base64Captcha.DefaultMemStore //默认的存储
var store base64Captcha.Store = RedisStore{}

//这里改成了redis作为存储captcha的存储器
//这是一个基于接口的设计，其实任何的keyvalue的存储器都可以设置为存储器，只要你实现了一个接口：Store接口，该接口的具体实现在redisStore中

//获取验证码
func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver             //driver就是一个driver,用来生成验证码的吧
	driverString := base64Captcha.DriverString{ //验证码的配置
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          3,
		Source:          "1234567890",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err

}
func VerifyCaptcha(id string, VerifyValue string) bool { //根据id和value去验证.直接调库
	fmt.Println("验证码id:", id, "验证值", VerifyValue) //获取不到id
	fmt.Println("验证码鉴定")
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
