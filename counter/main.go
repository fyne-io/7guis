package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := binding.NewInt()
	button := widget.NewButton("Count", func() {
		i, _ := count.Get()
		_ = count.Set(i+1)
	})

	w.SetContent(container.NewGridWithColumns(2,
		widget.NewLabelWithData(binding.IntToString(count)), button))
	w.ShowAndRun()
}
