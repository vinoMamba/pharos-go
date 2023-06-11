package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	return gomail.NewDialer(
		viper.GetString("EMAIL_HOST"),
		viper.GetInt("EMAIL_PORT"),
		viper.GetString("EMAIL_USER"),
		viper.GetString("EMAIL_PASSWD"),
	)
}

func newMessage(to, subject, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "1@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func Send() {
	m := newMessage("2@qq.com", "Hello!", "Hello <b>Bob</b> and <i>Cora</i>!")
	d := newDialer()
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendValidationCode(email, code string) error {
	m := newMessage(email, "VinoApp 的验证码", "您的验证码是："+code+"，5分钟内有效。")
	d := newDialer()
	return d.DialAndSend(m)
}
