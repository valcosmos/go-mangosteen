package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@admin.com")
	m.SetHeader("To", "admin@admin.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Cupid</b>")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
