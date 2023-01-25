// Package mail
// @Author lyf
// @Update lyf 2023.01
package mail

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"gopkg.in/gomail.v2"
	"regexp"
)

//******************************************************************//
//							email 结构体								//
//******************************************************************//

type email struct {
	entity *gomail.Dialer
}

// NewMessageManual
//
//	@description	手动构建邮件，若无特殊需求可以使用其他函数实现业务
//	@param	fromUserEmail	发送者邮箱
//	@param	fromUserName	发送者名称
//	@param	toUser	接受者邮箱
//	@param	needCopy	需要抄送则为真，否则为假
//	@param	copyUserName	需要抄送人的名称
//	@param	copyUserAddress	需要抄送人的邮箱地址
//	@param	title	邮箱标题
//	@param	body	邮箱内容
//	@param	needAttach 需要附件
//	@param	attach 附件路径
func (e *email) NewMessageManual(fromUserEmail string, fromUserName string, toUser string, needCopy bool,
	copyUserName string, copyUserAddress string,
	title string, body string, needAttach bool, attach string) error {

	if !VerifyEmailFormat(fromUserEmail) || !VerifyEmailFormat(toUser) {
		return errors.New(fmt.Sprintf("邮箱格式错误:fromUser:%s,toUser:%s", fromUserEmail, toUser))
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(fromUserEmail, fromUserName))
	m.SetHeader("To", toUser)
	if needCopy {
		if !VerifyEmailFormat(copyUserAddress) {
			return errors.New(fmt.Sprintf("邮箱格式错误:copyUserAddress:%s", copyUserAddress))
		}
		m.SetAddressHeader("Cc", copyUserAddress, copyUserName)
	}
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	if needAttach {
		m.Attach(attach)
	}

	return e.entity.DialAndSend(m)
}

// NewVerificationCode
//
//	@description	发送验证码，option为操作内容，如注册账号，code是验证码
func (e *email) NewVerificationCode(option string, toUser string, code string) error {

	title := fmt.Sprintf("XCPCBoard %s验证码", option)
	body := fmt.Sprintf(VerificationCodeBody, option, code)
	return e.NewMessageManual(config.Conf.Mail.UserName, "XCPCBoard", toUser,
		false, "", "", title, body, false, "")
}

// NewCreateUserMsg 给待注册的用户发送验证码
func (e *email) NewCreateUserMsg(toUser string, code string) error {

	return e.NewVerificationCode("注册账号", toUser, code, )
}

func (e *email) SentMsgToAdmin(msg string) error {
	//SentToAdminBody
	title := "XCPCBoard系统出错"
	body := fmt.Sprintf(SentToAdminBody, msg)
	return e.NewMessageManual(config.Conf.Mail.UserName, "XCPCBoard", config.Conf.Admin.Email,
		false, "", "", title, body, false, "")
}

//******************************************************************//
//							email 初始化								//
//******************************************************************//

var Email *email

func InitEmail() {
	m := gomail.NewDialer(config.Conf.Mail.Host, config.Conf.Mail.Port,
		config.Conf.Mail.UserName, config.Conf.Mail.Password)
	Email = new(email)
	Email.entity = m
}

// ******************************************************************//
//
//	util								//
//
// ******************************************************************//
const (
	//VerificationCodeBody 注册用户的HTML的body
	VerificationCodeBody = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <base target="_blank" />
    <style type="text/css">::-webkit-scrollbar{ display: none; }</style>
    <style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>
    <style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
    <style type="text/css">
        body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
        td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
        pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word}
        th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
        img{ border:0}
        header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
        blockquote{margin-right:0px}
    </style>
</head>
<body tabindex="0" role="listitem">
<table width="700" border="0" align="center" cellspacing="0" style="width:700px;">
    <tbody>
    <tr>
        <td>
            <div style="width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;">
                <table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;">
                    <tbody><tr><td width="210"></td></tr></tbody>
                </table>
            </div>
            <div style="width:680px;padding:0 10px;margin:0 auto;">
                <div style="line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;">
                    <strong style="display:block;margin-bottom:15px;">尊敬的用户：<span style="color:#f60;font-size: 16px;"></span>您好！</strong>
                    <strong style="display:block;margin-bottom:15px;">
                        您正在进行<span style="color: red">%s</span>操作，请在验证码输入框中输入：<span style="color:#f60;font-size: 24px">%s</span>，以完成操作。
                    </strong>
                </div>
                <div style="margin-bottom:30px;">
                    <small style="display:block;margin-bottom:20px;font-size:12px;">
                        <p style="color:#747474;">
                            注意：此操作可能会修改您的密码、登录邮箱或绑定手机。如非本人操作，请及时登录并修改密码以保证帐户安全
                            <br>（工作人员不会向你索取此验证码，请勿泄漏！)
                        </p>
                    </small>
                </div>
            </div>
            <div style="width:700px;margin:0 auto;">
                <div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
                    <p>此为系统邮件，请勿回复<br>
                        请保管好您的邮箱，避免账号被他人盗用
                    </p>
                    <p>XCPCBoard</p>
                </div>
            </div>
        </td>
    </tr>
    </tbody>
</table>
</body>
</html>


`

	//SentToAdminBody 注册用户的HTML的body
	SentToAdminBody = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <base target="_blank" />
    <style type="text/css">::-webkit-scrollbar{ display: none; }</style>
    <style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>
    <style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
    <style type="text/css">
        body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
        td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
        pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word}
        th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
        img{ border:0}
        header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
        blockquote{margin-right:0px}
    </style>
</head>
<body tabindex="0" role="listitem">
<table width="700" border="0" align="center" cellspacing="0" style="width:700px;">
    <tbody>
    <tr>
        <td>
            <div style="width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;">
                <table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;">
                    <tbody><tr><td width="210"></td></tr></tbody>
                </table>
            </div>
            <div style="width:680px;padding:0 10px;margin:0 auto;">
                <div style="line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;">
                    <strong style="display:block;margin-bottom:15px;">尊敬的用户：<span style="color:#f60;font-size: 16px;"></span>您好！</strong>
                    <strong style="display:block;margin-bottom:15px;">
                        您的XCPCBoard系统遇到严重错误，详情为:%s
                    </strong>
                </div>
                
            </div>
            <div style="width:700px;margin:0 auto;">
                <div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
                    <p>此为系统邮件，请勿回复<br>
                        请保管好您的邮箱，避免账号被他人盗用
                    </p>
                    <p>XCPCBoard</p>
                </div>
            </div>
        </td>
    </tr>
    </tbody>
</table>
</body>
</html>`
)

// VerifyEmailFormat
//
// email verify
func VerifyEmailFormat(email string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//// VerifyMobileFormat
////
//// mobile verify
//func VerifyMobileFormat(mobileNum string) bool {
//	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
//
//	reg := regexp.MustCompile(regular)
//	return reg.MatchString(mobileNum)
//}
