package client

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RenderClientView(a fyne.App, w fyne.Window) {

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Host Code")

	view := container.NewVBox(
		widget.NewButton("Generate Client Code", func() {

			widget.NewModalPopUp(widget.NewLabel("Client Code copied to clipboard"), w.Canvas()).Show()
			w.Clipboard().SetContent("--code--")
		}),
		input,
		widget.NewButton("Save", func() {
			log.Println("Content was:", input.Text)

			renderClientConnectionView(a, w)
		}))

	w.SetContent(view)

}
