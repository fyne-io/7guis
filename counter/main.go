package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dataapi"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := dataapi.NewInt(0)
	button := widget.NewButton("Count", func() {
		count.SetInt(count.Value()+1)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewLabel("0").Bind(count), button))
	w.ShowAndRun()
}
