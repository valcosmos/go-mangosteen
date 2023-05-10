package email

import "gopkg.in/gomail.v2"

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@admin.com")
	m.SetHeader("To", "admin@admin.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Cupid</b>")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.example.com", 456, "admin@admin.com", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
