package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/gomail.v2"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window with a title and size
	myWindow := myApp.NewWindow("Email Sender")
	myWindow.Resize(fyne.Size{Width: 600, Height: 500})

	// Create input fields for the email message
	fromInput := widget.NewEntry()
	fromInput.SetPlaceHolder("From")

	toInput := widget.NewEntry()
	toInput.SetPlaceHolder("To")

	subjectInput := widget.NewEntry()
	subjectInput.SetPlaceHolder("Subject")

	bodyInput := widget.NewMultiLineEntry()
	bodyInput.SetPlaceHolder("Write your message here :)")

	// Create a button to send the email
	sendButton := widget.NewButton("Send", func() {
		// Create a new email message
		m := gomail.NewMessage()

		// Set the sender and recipient email addresses
		m.SetHeader("From", fromInput.Text)
		m.SetHeader("To", toInput.Text)

		// Set the email subject and body
		m.SetHeader("Subject", subjectInput.Text)
		m.SetBody("text/plain", bodyInput.Text)

		// Create a new mailer object with the SMTP server configuration
		d := gomail.NewDialer("smtp.gmail.com", 587, "[Email Here]", "[App-password here]")

		// Send the email message
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Failed to send email:", err)
		} else {
			fmt.Println("Email sent successfully!")
		}
	})

	// Create a container to hold the input fields and button
	content := container.NewVBox(
		fromInput,
		toInput,
		subjectInput,
		bodyInput,
		layout.NewSpacer(),
		sendButton,
	)

	// Set the container as the window's content
	myWindow.SetContent(content)

	// Show the window and run the Fyne application
	myWindow.ShowAndRun()
}
