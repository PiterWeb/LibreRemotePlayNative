package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PiterWeb/LibreRemotePlayNative/views/client"
)

func main() {
	a := app.New()
	w := a.NewWindow("LibreRemotePlay")
	w.Resize(fyne.Size{Width: 1280, Height: 720})

	clientBtn := widget.NewButton("Client", func() {

		log.Println("Btn tapped")

		client.RenderClientView(a,w)

	})

	clientBtn.Alignment = widget.ButtonAlignCenter

	hostBtn := widget.NewButton("Host", func() {

		log.Println("Btn tapped")

	})

	hostBtn.Alignment = widget.ButtonAlignCenter

	view := container.NewGridWithColumns(2, hostBtn, clientBtn)

	w.SetContent(view)

	w.ShowAndRun()
}
