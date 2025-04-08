package main

import (
	"image"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{Width: 1280, Height: 720})

	img, err := getImageFromFilePath("./example.png")

	if err != nil {
		return
	}

	view := container.NewAdaptiveGrid(1, canvas.NewImageFromImage(img))

	w.SetContent(view)

	w.ShowAndRun()
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
