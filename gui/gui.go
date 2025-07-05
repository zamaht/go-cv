package gui

import (
	"basic/colorutil"
	"basic/imageutil"
	"fmt"
	"image"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

func addBalanceSlider(a *fyne.App, img *image.Image, imageWindow *fyne.Window) {
	slider := widget.NewSlider(1, 30)

	slider.OnChanged = func(balance float64) {
		processedImage := colorutil.BalanceColor(*img, uint32(balance))
		newCanvasImage := canvas.NewImageFromImage(processedImage)
		(*imageWindow).SetContent(newCanvasImage)
	}

	sliderWindow := (*a).NewWindow("balance")
	sliderWindow.Resize(fyne.Size{Width: 200, Height: 100})
	sliderWindow.SetContent(slider)
	sliderWindow.Show()
}

func ProcessImage(imageName string) {
	a := app.New()

	imageWindow := a.NewWindow(imageName)

	img := imageutil.ReadImage(imageName)
	fmt.Printf("img.ColorModel(): %v\n", img.ColorModel())

	canvasImage := canvas.NewImageFromImage(img)
	fmt.Printf("canvasImage: %v\n", canvasImage)
	canvasImage.FillMode = canvas.ImageFillContain

	imageWindow.Resize(fyne.Size{Width: 640, Height: 480})
	imageWindow.SetContent(canvasImage)
	imageWindow.Show()

	addBalanceSlider(&a, &canvasImage.Image, &imageWindow)

	a.Run()
}
