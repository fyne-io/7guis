package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	update, delete *widget.Button
	name, surname  *widget.Entry
	list           *widget.List

	filtered []int
	selected int
}

func (g *gui) createDelete() *widget.Button {
	var btn *widget.Button
	btn = widget.NewButton("Delete", func() {
		if g.selected < 0 || g.selected >= len(people) || len(people) == 0 {
			return
		}

		if g.selected == 0 {
			people = people[1:]
		} else if g.selected == len(people)-1 {
			people = people[:len(people)-1]
		} else {
			people = append(people[:g.selected], people[g.selected+1:]...)
		}
		g.filtered = noFilter()
		g.list.UnselectAll()
		g.list.Refresh()
		g.update.Disable()
		btn.Disable()
	})
	btn.Disable()
	return btn
}

func (g *gui) createFilter() *widget.Entry {
	f := widget.NewEntry()
	f.OnChanged = func(prefix string) {
		g.list.UnselectAll()
		g.update.Disable()
		g.delete.Disable()
		if prefix == "" {
			g.filtered = noFilter()
			g.list.Refresh()
			return
		}

		prefix = strings.ToLower(prefix)
		f := []int{}
		for i, p := range people {
			if strings.Index(strings.ToLower(p.surname), prefix) == 0 {
				f = append(f, i)
			}
		}
		g.filtered = f
		g.list.Refresh()
	}
	return f
}

func (g *gui) createList() *widget.List {
	l := widget.NewList(func() int {
		return len(g.filtered)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(people[g.filtered[id]].String())
	})

	l.OnSelected = func(id widget.ListItemID) {
		g.selected = g.filtered[id]
		g.name.SetText(people[g.selected].name)
		g.surname.SetText(people[g.selected].surname)

		g.update.Enable()
		g.delete.Enable()
	}
	l.OnUnselected = func(id widget.ListItemID) {
		g.update.Disable()
		g.delete.Disable()
	}

	return l
}

func (g *gui) createNew() *widget.Button {
	return widget.NewButton("Create", func() {
		p := &person{name: g.name.Text, surname: g.surname.Text}
		people = append(people, p)
		g.filtered = noFilter()
		g.list.Refresh()
		g.list.Select(len(people) - 1)
	})
}

func (g *gui) createUpdate() *widget.Button {
	btn := widget.NewButton("Update", func() {
		if g.selected < 0 || g.selected >= len(people) {
			return
		}

		people[g.selected].name = g.name.Text
		people[g.selected].surname = g.surname.Text
		g.list.Refresh()
	})
	btn.Disable()
	return btn
}

func main() {
	a := app.New()
	w := a.NewWindow("CRUD")

	g := gui{name: widget.NewEntry(), surname: widget.NewEntry(), filtered: noFilter(), selected: -1}
	g.list = g.createList()
	g.update = g.createUpdate()
	g.delete = g.createDelete()

	form := widget.NewForm(
		widget.NewFormItem("Name", g.name),
		widget.NewFormItem("Surname", g.surname))

	top := container.NewGridWithColumns(2,
		widget.NewForm(widget.NewFormItem("Filter prefix:", g.createFilter())))
	bottom := container.NewHBox(g.createNew(),
		g.update, g.delete)

	grid := container.NewGridWithColumns(2, g.list, form)
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
