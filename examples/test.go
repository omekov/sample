package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	// user we are authorizing as
	from := "omekovazamat@gmail.com"

	// use we are sending email to
	to := "umekovazamat@gmail.com"

	// server we are authorized to send email through
	host := "gmail.com"

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", from, "", host)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	message := `To: "Some User" <someuser@example.com>
From: "Other User" <otheruser@example.com>
Subject: Testing Email From Go!!

This is the message we are sending. That's it!
`

	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")
}
