package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Email Interface")

	w.Resize(fyne.NewSize(800, 600))

	// Create text boxes for email fields
	toEntry := widget.NewEntry()
	toEntry.SetPlaceHolder("To")

	ccEntry := widget.NewEntry()
	ccEntry.SetPlaceHolder("CC")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("Subject")

	bodyEntry := widget.NewMultiLineEntry()

	// Create buttons
	sendButton := widget.NewButton("Send", func() {
		// Send email functionality here
	})

	cancelButton := widget.NewButton("Cancel", func() {
		// Close window functionality here
		w.Close()
	})

	// Create container for buttons
	buttonBox := container.NewHBox(sendButton, cancelButton)

	// Create container for text boxes
	entryBox := container.NewVBox(toEntry, ccEntry, subjectEntry, bodyEntry)

	// Create container for all elements
	content := container.NewVBox(entryBox, buttonBox)

	// Set the window content and show the window
	w.SetContent(content)
	w.ShowAndRun()
}
