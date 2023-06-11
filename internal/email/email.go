package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1@qq.com")
	m.SetHeader("To", "2@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(
		viper.GetString("EMAIL_HOST"),
		viper.GetInt("EMAIL_PORT"),
		viper.GetString("EMAIL_USER"),
		viper.GetString("EMAIL_PASSWD"),
	)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
