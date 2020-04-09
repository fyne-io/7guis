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

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	// Create binding
	temperature := NewTemperatureBinding()

	// Create widgets
	inputC := widget.NewEntry().BindText(temperature.Celsius)
	inputF := widget.NewEntry().BindText(temperature.Fahrenheit)

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}

// TemperatureBinding implements a data binding for temperature in Kelvin.
type TemperatureBinding struct {
	binding.Float64Binding
	Celsius, Fahrenheit *binding.StringBinding
}

// NewTemperatureBinding creates and returns a new TemperatureBinding.
func NewTemperatureBinding() *TemperatureBinding {
	kb := &TemperatureBinding{
		Celsius:    binding.NewStringBinding(""),
		Fahrenheit: binding.NewStringBinding(""),
	}
	kb.Celsius.AddListener(parseFloat64To(kb.SetCelsius))
	kb.Fahrenheit.AddListener(parseFloat64To(kb.SetFahrenheit))
	kb.AddListener(func(float64) {
		kb.Celsius.Set(fmt.Sprintf("%.2f", kb.GetCelsius()))
		kb.Fahrenheit.Set(fmt.Sprintf("%.2f", kb.GetFahrenheit()))
	})
	return kb
}

// GetCelsius returns the temperature converted to Celsius.
func (b *TemperatureBinding) GetCelsius() float64 {
	return b.Get() - 273.15
}

// GetFahrenheit returns the temperature converted to Fahrenheit.
func (b *TemperatureBinding) GetFahrenheit() float64 {
	return (b.Get()-273.15)*9.0/5.0 + 32
}

// SetCelsius updates the temperature to the given reading in Celsius.
func (b *TemperatureBinding) SetCelsius(c float64) {
	b.Set(c + 273.15)
}

// SetFahrenheit updates the temperature to the given reading in Fahrenheit.
func (b *TemperatureBinding) SetFahrenheit(f float64) {
	b.Set((f-32)*5.0/9.0 + 273.15)
}

func parseFloat64To(c func(float64)) func(string) {
	return func(s string) {
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			c(f)
		}
	}
}
