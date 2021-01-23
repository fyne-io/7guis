package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	valueC := binding.NewFloat()
	valueF := celsiusToFarenheit(valueC)

	entryC := widget.NewEntryWithData(binding.FloatToString(valueC))
	entryC.Wrapping = fyne.TextWrapOff
	entryF := widget.NewEntryWithData(binding.FloatToString(valueF))
	entryF.Wrapping = fyne.TextWrapOff
	w.SetContent(container.NewGridWithColumns(4,
		entryC, widget.NewLabel("Celsius ="),
		entryF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}

type cToF struct {
	binding.Float
}

func (c *cToF) Get() (float64, error) {
	cDeg, _ := c.Float.Get()
	fDeg := cDeg*(9.0/5.0) + 32
	return fDeg, nil
}

func (c *cToF) Set(f float64) error {
	cDeg := (f - 32) * (5.0 / 9.0)
	_ = c.Float.Set(cDeg)
	return nil
}

func celsiusToFarenheit(in binding.Float) binding.Float {
	return &cToF{in}
}
