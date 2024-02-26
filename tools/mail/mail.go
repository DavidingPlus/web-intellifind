package mail

import (
	"github.com/spf13/viper"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

// NewMailer 单例模式获取
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	config := map[string]string{
		"smtp_username": viper.GetString("Mail.smtp.username"),
		"smtp_password": viper.GetString("Mail.smtp.password"),
		"smtp_host":     viper.GetString("Mail.smtp.host"),
		"smtp_port":     viper.GetString("Mail.smtp.port"),
	}
	return mailer.Driver.Send(email, config)
}
