package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	selected := -1
	a := app.New()
	w := a.NewWindow("CRUD")

	name := widget.NewEntry()
	surname := widget.NewEntry()
	list := widget.NewList(func() int {
		return len(people)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(people[id].String())
	})

	update := widget.NewButton("Update", func() {
		if selected < 0 || selected >= len(people) {
			return
		}

		people[selected].name = name.Text
		people[selected].surname = surname.Text
		list.Refresh()
	})
	update.Disable()
	delete := widget.NewButton("Delete", func() {
			if selected < 0 || selected >= len(people) || len(people) == 0 {
				return
			}

			if selected == 0 {
				people = people[1:]
			} else if selected == len(people) - 1 {
				people = people[:len(people)-1]
			} else {
				people = append(people[:selected], people[selected+1:]...)
			}
			list.Refresh()
		})
	delete.Disable()

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
		name.SetText(people[id].name)
		surname.SetText(people[id].surname)

		update.Enable()
		delete.Enable()
	}
	list.OnUnselected = func(id widget.ListItemID) {
		update.Disable()
		delete.Disable()
	}

	form := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Surname", surname))

	top := container.NewGridWithColumns(2,
		widget.NewForm(widget.NewFormItem("Filter prefix:", widget.NewEntry())))
	bottom := container.NewHBox(
		widget.NewButton("Create", func() {
			p := &person{name: name.Text, surname: surname.Text}
			people = append(people, p)
			list.Refresh()
			list.Select(len(people)-1)
		}),
		update, delete)

	grid := container.NewGridWithColumns(2, list, form)
	w.SetContent(
		container.NewBorder(top, bottom, nil, nil, grid))
	w.ShowAndRun()
}
