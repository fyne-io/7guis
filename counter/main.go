package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	// Create bindings
	bindInt := &binding.IntBinding{}
	bindString := &binding.StringBinding{}

	// Configure int to string pipeline
	bindInt.AddListener(func(i int) {
		bindString.Set(fmt.Sprintf("%d", i))
	})

	// Create widgets
	value := widget.NewLabel("0")
	value.BindText(bindString)
	button := widget.NewButton("Count", func() {
		bindInt.Set(bindInt.Get() + 1)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		value, button))
	w.ShowAndRun()
}
