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

	// Create binding
	bindingKelvin := NewKelvinBinding()

	// Create widgets
	inputC := widget.NewEntry()
	inputC.BindText(bindingKelvin.Celsius)
	inputF := widget.NewEntry()
	inputF.BindText(bindingKelvin.Fahrenheit)

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}

type KelvinBinding struct {
	binding.Float64Binding
	Celsius, Fahrenheit *binding.StringBinding
}

func NewKelvinBinding() *KelvinBinding {
	kb := &KelvinBinding{
		Celsius:    &binding.StringBinding{},
		Fahrenheit: &binding.StringBinding{},
	}
	kb.Celsius.AddListener(func(text string) {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		kb.SetCelsius(f)
	})
	kb.Fahrenheit.AddListener(func(text string) {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		kb.SetFahrenheit(f)
	})
	kb.AddListener(func(float64) {
		kb.Celsius.Set(fmt.Sprintf("%.2f", kb.GetCelsius()))
		kb.Fahrenheit.Set(fmt.Sprintf("%.2f", kb.GetFahrenheit()))
	})
	return kb
}

func (b *KelvinBinding) GetCelsius() float64 {
	return b.Get() - 273.15
}

func (b *KelvinBinding) GetFahrenheit() float64 {
	return (b.Get()-273.15)*9.0/5.0 + 32
}

func (b *KelvinBinding) SetCelsius(c float64) {
	b.Set(c + 273.15)
}

func (b *KelvinBinding) SetFahrenheit(f float64) {
	b.Set((f-32)*5.0/9.0 + 273.15)
}
