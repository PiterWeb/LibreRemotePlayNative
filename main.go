package main

import (
	"bufio"
	"bytes"
	"image"
	"image/jpeg"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{Width: 1280, Height: 720})

	imgReader := new(bufio.Reader)
	imgWriter := new(bufio.Writer)

	imgReaderWriter := bufio.NewReadWriter(imgReader, imgWriter)

	img, err := os.ReadFile("./example.png")

	if err != nil {
		return
	}

	_, _ = imgReaderWriter.Write(img)

	view := container.NewAdaptiveGrid(1, canvas.NewImageFromReader(imgReaderWriter, "main_image"))

	w.SetContent(view)

	go updateView(view, imgWriter)

	w.ShowAndRun()
}

func updateView(view *fyne.Container, imgWriter *bufio.Writer) {

	if canvasImg, ok := view.Objects[0].(*canvas.Image); ok {
		for range time.Tick(time.Second) {

			if canvasImg.Image == image.Black {
				ex_img, err := os.ReadFile("./example.png")

				if err != nil {
					continue
				}

				_, _ = imgWriter.Write(ex_img)
			} else {
				buf := new(bytes.Buffer)
				err := jpeg.Encode(buf, image.Black, nil)

				if err != nil {
					continue
				}

				_, _ = imgWriter.Write(buf.Bytes())
			}

			canvasImg.Refresh()

		}
	}
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
