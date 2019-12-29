package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var count int

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	value := widget.NewLabel("0")
	button := widget.NewButton("Count", func() {
		count++
		value.SetText(fmt.Sprintf("%d", count))
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		value, button))
	w.ShowAndRun()
}
