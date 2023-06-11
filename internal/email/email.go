package email

import (
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1@qq.com")
	m.SetHeader("To", "2@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("0.0.0.0", 1025, "", "")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
