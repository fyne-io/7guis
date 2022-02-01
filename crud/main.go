package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var filtered = noFilter()

func main() {
	selected := -1
	a := app.New()
	w := a.NewWindow("CRUD")

	name := widget.NewEntry()
	surname := widget.NewEntry()
	list := widget.NewList(func() int {
		return len(filtered)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(people[filtered[id]].String())
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
	var delete *widget.Button
	delete = widget.NewButton("Delete", func() {
		if selected < 0 || selected >= len(people) || len(people) == 0 {
			return
		}

		if selected == 0 {
			people = people[1:]
		} else if selected == len(people)-1 {
			people = people[:len(people)-1]
		} else {
			people = append(people[:selected], people[selected+1:]...)
		}
		filtered = noFilter()
		list.UnselectAll()
		list.Refresh()
		update.Disable()
		delete.Disable()
	})
	delete.Disable()

	list.OnSelected = func(id widget.ListItemID) {
		selected = filtered[id]
		name.SetText(people[selected].name)
		surname.SetText(people[selected].surname)

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

	filter := widget.NewEntry()
	filter.OnChanged = func(prefix string) {
		list.UnselectAll()
		update.Disable()
		delete.Disable()
		if prefix == "" {
			filtered = noFilter()
			list.Refresh()
			return
		}

		prefix = strings.ToLower(prefix)
		f := []int{}
		for i, p := range people {
			if strings.Index(strings.ToLower(p.surname), prefix) == 0 {
				f = append(f, i)
			}
		}
		filtered = f
		list.Refresh()
	}
	top := container.NewGridWithColumns(2,
		widget.NewForm(widget.NewFormItem("Filter prefix:", filter)))
	bottom := container.NewHBox(
		widget.NewButton("Create", func() {
			p := &person{name: name.Text, surname: surname.Text}
			people = append(people, p)
			filtered = noFilter()
			list.Refresh()
			list.Select(len(people) - 1)
		}),
		update, delete)

	grid := container.NewGridWithColumns(2, list, form)
	w.SetContent(
		container.NewBorder(top, bottom, nil, nil, grid))
	w.ShowAndRun()
}

func noFilter() []int {
	all := make([]int, len(people))
	for i := range people {
		all[i] = i
	}
	return all
}
