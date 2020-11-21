package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/data/binding"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	valueC := binding.NewFloat()
	valueF := binding.NewFloat()
	valueC.AddListener(binding.NewDataListener(func() {
		fDeg := valueC.Get()*(9.0/5.0) + 32
		valueF.Set(fDeg)
	}))
	valueF.AddListener(binding.NewDataListener(func() {
		cDeg := (valueF.Get() - 32) * (5.0 / 9.0)
		valueC.Set(cDeg)
	}))

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		widget.NewEntryWithData(binding.FloatToString(valueC)), widget.NewLabel("Celsius ="),
		widget.NewEntryWithData(binding.FloatToString(valueF)), widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}
