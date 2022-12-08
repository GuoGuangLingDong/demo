package utils

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"github.com/mojocn/base64Captcha"
)

// var result = base64Captcha.DefaultMemStore
var result = base64Captcha.NewMemoryStore(1024, 1*time.Minute)

type imageCode struct {
}

var ImageCode = new(imageCode)

// stringConfig 生成图形化字符串验证码配置
func (r imageCode) stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           300,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          "ABCDEFGHIKMNPQRSTUVWXYZabcdefghikmnpqrstuvwxyz23456789",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return stringType
}

// CreateCode 生成验证码图片
func (r imageCode) CreateCode() (id, b64s string, err error) {
	driver := r.stringConfig()
	// 创建验证码并传入创建的类型的配置，以及存储的对象
	c := base64Captcha.NewCaptcha(driver, result)
	id, b64s, answer, err := r.Generate(c)
	fmt.Println("图形验证码", answer)
	if err != nil {
		return
	}
	// g.Log().Info("id:" + id + ",answer:" + answer)
	err = ImageCodeRedis.Set(id, answer)
	if err != nil {
		return
	}
	return id, b64s, err
}

// VerifyCaptcha 校验验证码
func (r imageCode) VerifyCaptcha(id, VerifyValue string) (err error) {
	// return result.Verify(id, VerifyValue, true)
	ret, err := ImageCodeRedis.Get(id, false)
	if err != nil {
		return
	}
	if ret == "" {
		err = fmt.Errorf("校验码过期 请重新获取!")
		return
	}
	// g.Log().Info("ret:" + strings.ToLower(ret))
	// g.Log().Info("value:" + strings.ToLower(VerifyValue))
	fmt.Println(ret)
	if strings.ToLower(ret) != strings.ToLower(VerifyValue) {
		// err = gerror.New("请输入正确的校验码:" + ret + "!")
		err = fmt.Errorf("请输入正确的校验码!")
		return
	}
	return
}

//Generate generates a random id, base64 image string or an error if any
func (r imageCode) Generate(c *base64Captcha.Captcha) (id, b64s, answer string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", answer, err
	}
	// err = c.Store.Set(id, answer)
	// if err != nil {
	// 	return "", "", answer, err
	// }
	b64s = item.EncodeB64string()
	return
}
