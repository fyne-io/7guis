package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const TIME_LAYOUT = "02.01.2006"

func main() {
	a := app.New()
	w := a.NewWindow("Flight Booker")

	flightTypes := []string{"one-way flight", "return flight"}

	// Create bindings
	flightType := &binding.StringBinding{}
	startDate := &binding.StringBinding{}
	returnDate := &binding.StringBinding{}

	// Create widgets
	combo := widget.NewSelect(flightTypes, nil)
	combo.BindSelected(flightType)
	entry1 := widget.NewEntry()
	entry1.BindText(startDate)
	entry2 := widget.NewEntry()
	entry2.BindText(returnDate)
	button := widget.NewButton("Book", func() {
		dialog.ShowInformation("Information", fmt.Sprintf("You booked a %s on %s", flightType.Get(), startDate.Get()), w)
	})

	// Configure behaviour
	update := func(string) {
		// Check entry1 is valid
		start, err := time.Parse(TIME_LAYOUT, startDate.Get())
		if err != nil {
			// TODO set entry1.background red
			button.Disable()
			return
		}
		if flightType.Get() == flightTypes[1] {
			// Check entry2 is valid
			ret, err := time.Parse(TIME_LAYOUT, returnDate.Get())
			if err != nil {
				// TODO set entry2.background red
				button.Disable()
				return
			}
			// If start is not before return, set button disabled
			if !start.Before(ret) {
				button.Disable()
				return
			}
		}
		button.Enable()
	}
	startDate.AddListener(update)
	returnDate.AddListener(update)
	flightType.AddListener(func(s string) {
		if s == flightTypes[1] {
			entry2.Enable()
		} else {
			entry2.Disable()
		}
		update(s)
	})

	// Set initial state
	flightType.Set(flightTypes[0])
	n := time.Now().Format(TIME_LAYOUT)
	startDate.Set(n)
	returnDate.Set(n)

	w.SetContent(widget.NewVBox(combo, entry1, entry2, button))
	w.ShowAndRun()
}
