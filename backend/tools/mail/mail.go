package mail

import "github.com/spf13/viper"

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

func (mailer *Mailer) Send(email Email) bool {
	config := map[string]string{
		"smtp_username": viper.GetString("mail.smtp_username"),
		"smtp_password": viper.GetString("mail.smtp_password"),
		"smtp_host":     viper.GetString("mail.smtp_host"),
		"smtp_port":     viper.GetString("mail.smtp_port"),
	}
	return mailer.Driver.Send(email, config)
}
