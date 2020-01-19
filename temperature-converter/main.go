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
	w := a.NewWindow("Temperature Converter")

	valueC := dataapi.NewFloat(0)
	valueF := dataapi.NewFloat(0)
	valueC.AddListener(func(data dataapi.DataItem) {
		fDeg := valueC.Value()*(9.0/5.0) + 32
		valueF.SetFloat(fDeg)
	})
	valueF.AddListener(func(data dataapi.DataItem) {
		cDeg := (valueF.Value() - 32) * (5.0 / 9.0)
		valueC.SetFloat(cDeg)
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		widget.NewEntry().Bind(valueC), widget.NewLabel("Celsius ="),
		widget.NewEntry().Bind(valueF), widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}
