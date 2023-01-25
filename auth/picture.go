// Package auth
// @Author lyf
// @Update lyf 2023.01
package auth

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 使用DefaultMemStore 创建的验证码存储对象，存储的验证码为 10240 个，过期时间为 10分钟
var authStorage = base64Captcha.DefaultMemStore

// stringConfig 生成图形化字符串验证码配置
func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           300,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          6,
		Source:          "123456789qwertyuiopasdfghjklzxcvbQWERTYUIOPASDFGHJKLZXCVBNM",
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

// CreateCode 创建图片验证码
//
//	@result id 验证码id
//	@result bse64s 图片base64编码
//	@result err 错误
func CreateCode() (string, string, error) {

	driver := stringConfig()

	// 创建验证码并传入创建的类型的配置，以及存储的对象
	c := base64Captcha.NewCaptcha(driver, authStorage)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 效验验证码
//
//	当为 true 时，校验 传入的id 的验证码，校验完 这个ID的验证码就要在内存中删除
//	当为 false 时，校验 传入的id 的验证码，校验完 这个ID的验证码不删除
//	@pram id 验证码id
//	@pram VerifyValue 用户输入的答案
//	@result true：正确，false：失败
func VerifyCaptcha(id, VerifyValue string) bool {
	// result 为步骤1 创建的图片验证码存储对象
	return authStorage.Verify(id, VerifyValue, true)
}

// GetCodeAnswer 获取验证码答案
//
//	@Pram codeId 验证码id
//	@Result 验证码答案
func GetCodeAnswer(codeId string) string {
	// result 为步骤1 创建的图片验证码存储对象
	return authStorage.Get(codeId, false)
}
