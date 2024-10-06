package main

import (
	"log"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	buttons   []fyne.CanvasObject
	labelName = []string{"7", "8", "9", "/",
		"4", "5", "6", "x",
		"1", "2", "3", "-",
		"C", "0", ".", "+"}
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")
	myWindow.Resize(fyne.NewSize(350, 500))
	myWindow.SetFixedSize(true)

	label := widget.NewLabel("0")
	label.Alignment = fyne.TextAlignTrailing
	label.TextStyle = fyne.TextStyle{Italic: true}

	for _, l := range labelName {
		btn := widget.NewButton(l, func() {
			if label.Text == "0" {
				label.SetText(l)
			} else {
				label.SetText(label.Text + l)
			}
			log.Println(l)
		})
		buttons = append(buttons, btn)
	}
	buttonsGrid := container.New(layout.NewGridLayout(4), buttons...)
	layout := container.NewVBox(layout.NewSpacer(), label, buttonsGrid)
	layout.Resize(fyne.NewSize(330, 500))

	myWindow.SetContent(layout)
	myWindow.ShowAndRun()
}
