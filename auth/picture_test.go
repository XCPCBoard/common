package auth

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

// TestPicture 测试图像验证码
func TestPicture(t *testing.T) {
	id, bs, err := CreateCode()
	if err != nil {
		t.Fatal("无法构建验证码")
	}
	//fmt.Println(bs)

	ddd, _ := base64.StdEncoding.DecodeString(bs[22:])
	err = ioutil.WriteFile("./testOutput.png", ddd, 0666)
	if err != nil {
		t.Fatal("无法生成图像")
	}

	qq := GetCodeAnswer(id)
	//fmt.Println(qq)
	if !VerifyCaptcha(id, qq) {
		t.Error("验证码错误")
	}
}
