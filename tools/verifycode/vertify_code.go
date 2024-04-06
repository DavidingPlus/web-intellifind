package verifycode

import (
	"backend/logger"
	"backend/tools/helpers"
	"backend/tools/mail"
	"backend/utils"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: utils.Redis,
			},
		}
	})

	return internalVerifyCode
}

// SendEmail 发送邮件验证码，调用示例：
//
//	verifycode.NewVerifyCode().SendEmail(request.Email)
func (vc *VerifyCode) SendEmail(email string) error {

	// 生成验证码
	code := vc.generateVerifyCode(email)

	content := fmt.Sprintf("<div\n        style=\"border: 1px solid #f4f4f4; border-top: 5px solid #9373ee; box-sizing: border-box;  width: 622px; text-align: center; margin: 30px auto;\">\n        <div style=\"margin: 80px auto 32px; font-size: 32px; font-weight: 600; line-height: 45px; letter-spacing: 0px;\">\n            您的邮箱验证码\n        </div>\n        <div style=\"font-size: 18px; font-weight: 400;\">\n            您正在进行邮箱验证操作，验证码为:\n        </div>\n        <div\n            style=\"width: 500px; background-color: rgba(147, 115, 238, 0.04); font-weight: 650; font-size: 35px; color: #9373EE; margin: 30px auto; padding: 20px; letter-spacing: 12px;\">\n            %v\n        </div>\n        <div\n            style=\"font-size: 14px; line-height: 1.6; color: #5c5c5c; padding: 5px 0; font-family: lucida Grande, Verdana, Microsoft YaHei;\">\n            此验证码 5 分钟内有效\n        </div>\n        <div\n            style=\"font-size: 14px; line-height: 1.6; color: #5c5c5c; padding: 5px 0; font-family: lucida Grande, Verdana, Microsoft YaHei;\">\n            如非本人操作\n        </div>\n        <div\n            style=\"font-size: 14px; line-height: 1.6; color: #5c5c5c; padding: 5px 0; font-family: lucida Grande, Verdana, Microsoft YaHei;\">\n            转给他人将导致账号被盗和个人信息泄漏，谨防诈骗\n        </div>\n        <div style=\"height: 110px;\"></div>\n        <div style=\"color: #9373ee; text-decoration: none;\">\n            <a href=\"http://8.137.100.0/\">Web智寻</a> - 节省研发团队的每一分钟\n        </div>\n    </div>", code)
	// 发送邮件
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: viper.GetString("Mail.from.address"),
			Name:    viper.GetString("Mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) generateVerifyCode(key string) string {

	// 生成随机码
	code := helpers.RandomNumber(viper.GetInt("Verifycode.code_length"))

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
