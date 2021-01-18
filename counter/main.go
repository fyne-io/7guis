package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := binding.NewInt()
	button := widget.NewButton("Count", func() {
		i, _ := count.Get()
		count.Set(i+1)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewLabelWithData(binding.IntToString(count)), button))
	w.ShowAndRun()
}
