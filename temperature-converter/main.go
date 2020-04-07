package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	// Create bindings
	bindingFloat64C := &binding.Float64Binding{}
	bindingFloat64F := &binding.Float64Binding{}
	bindingStringC := &binding.StringBinding{}
	bindingStringF := &binding.StringBinding{}

	// Create widgets
	inputC := widget.NewEntry()
	inputC.BindText(bindingStringC)
	inputF := widget.NewEntry()
	inputF.BindText(bindingStringF)

	// Configure pipeline
	bindingStringC.AddListener(func(text string) {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		bindingFloat64F.Set(f*(9.0/5.0) + 32)
	})

	bindingStringF.AddListener(func(text string) {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		bindingFloat64C.Set((f - 32) * (5.0 / 9.0))
	})
	bindingFloat64C.AddListener(func (f float64) {
		bindingStringC.Set(fmt.Sprintf("%f", f))
	})
	bindingFloat64F.AddListener(func (f float64) {
		bindingStringF.Set(fmt.Sprintf("%f", f))
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}
