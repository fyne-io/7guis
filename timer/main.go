package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Timer")

	// Create bindings
	bindDuration := &binding.Float64Binding{}
	bindElapsed := &binding.Float64Binding{}
	bindElapsedString := &binding.StringBinding{}

	// Create widgets
	elapsedProgressBar := widget.NewProgressBar().BindValue(bindElapsed).BindMax(bindDuration)
	elapsedLabel := widget.NewLabel("").BindText(bindElapsedString)
	durationSlider := widget.NewSlider(1, 100).BindValue(bindDuration)
	resetButton := widget.NewButton("Reset", func() {
		bindElapsed.Set(0.0)
	})

	// Configure pipeline
	bindElapsed.AddListener(func(f float64) {
		bindElapsedString.Set(fmt.Sprintf("%.1fs", f))
	})

	// Create timer to increment elapsed time
	done := make(chan bool)
	defer func() {
		done <- true
	}()
	timer := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-timer.C:
				d := bindDuration.Get()
				e := bindElapsed.Get()
				if e < d {
					bindElapsed.Set(e + 1)
				}
			}
		}
	}()

	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewLabel("Elapsed Time:"), elapsedProgressBar,
		layout.NewSpacer(), elapsedLabel,
		widget.NewLabel("Duration:"), durationSlider,
		layout.NewSpacer(), resetButton))
	w.ShowAndRun()
}
