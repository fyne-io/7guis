package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
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

	inputC := widget.NewEntry()
	inputF := widget.NewEntry()

	inputC.OnChanged = func(text string) {
		if !isNumeric(inputC.Text) {
			return
		}

		cDeg, _ := strconv.Atoi(inputC.Text)
		fDeg := float64(cDeg)*(9.0/5.0) + 32

		inputF.Text = strconv.Itoa(int(fDeg))
		inputF.Refresh()
	}

	inputF.OnChanged = func(text string) {
		if !isNumeric(inputF.Text) {
			return
		}

		fDeg, _ := strconv.Atoi(inputF.Text)
		cDeg := (float64(fDeg) - 32) * (5.0 / 9.0)

		inputC.Text = strconv.Itoa(int(cDeg))
		inputC.Refresh()
	}

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}
