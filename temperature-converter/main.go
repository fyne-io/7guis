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
	bindingKelvin := &KelvinBinding{}
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
		bindingKelvin.SetCelsius(f)
	})

	bindingStringF.AddListener(func(text string) {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		bindingKelvin.SetFahrenheit(f)
	})
	bindingKelvin.AddListener(func (f float64) {
		bindingStringC.Set(fmt.Sprintf("%.2f", bindingKelvin.GetCelsius()))
		bindingStringF.Set(fmt.Sprintf("%.2f", bindingKelvin.GetFahrenheit()))
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}

type KelvinBinding struct {
	binding.Float64Binding
}

func (b *KelvinBinding) GetCelsius() float64 {
	return b.Get()-273.15
}

func (b *KelvinBinding) GetFahrenheit() float64 {
	return (b.Get() - 273.15) * 9.0/5.0 + 32
}

func (b *KelvinBinding) SetCelsius(c float64) {
	b.Set(c+273.15)
}

func (b *KelvinBinding) SetFahrenheit(f float64) {
	b.Set((f - 32) * 5.0/9.0 + 273.15)
}
