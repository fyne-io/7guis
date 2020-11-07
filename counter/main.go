package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := binding.NewInt()
	button := widget.NewButton("Count", func() {
		count.Set(count.Get()+1)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewLabelWithData(binding.IntToString(count)), button))
	w.ShowAndRun()
}
