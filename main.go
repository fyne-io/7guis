package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("CRUD")

	list := widget.NewList(func() int {
		return len(people)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(people[id].String())
	})

	form := widget.NewForm(
		widget.NewFormItem("Name", widget.NewEntry()),
		widget.NewFormItem("Surname", widget.NewEntry()))

	top := container.NewGridWithColumns(2,
		widget.NewForm(widget.NewFormItem("Filter prefix:", widget.NewEntry())))
	bottom := container.NewHBox(
		widget.NewButton("Create", func() {}),
		widget.NewButton("Update", func() {}),
		widget.NewButton("Delete", func() {}))

	grid := container.NewGridWithColumns(2, list, form)
	w.SetContent(
		container.NewBorder(top, bottom, nil, nil, grid))
	w.ShowAndRun()
}
