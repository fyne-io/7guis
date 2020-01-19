package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var humanInteraction bool = true

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
		if !isNumeric(inputC.Text) || !humanInteraction {
			return
		}

		cDeg, _ := strconv.Atoi(inputC.Text)
		fDeg := float64(cDeg)*(9.0/5.0) + 32

		humanInteraction = false
		inputF.SetText(strconv.Itoa(int(fDeg)))
		humanInteraction = true
	}

	inputF.OnChanged = func(text string) {
		if !isNumeric(inputF.Text) || !humanInteraction {
			return
		}

		fDeg, _ := strconv.Atoi(inputF.Text)
		cDeg := (float64(fDeg) - 32) * (5.0 / 9.0)

		humanInteraction = false
		inputC.SetText(strconv.Itoa(int(cDeg)))
		humanInteraction = true
	}

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		inputC, widget.NewLabel("Celsius ="), inputF, widget.NewLabel("Fahrenheit")))

	w.ShowAndRun()
}
