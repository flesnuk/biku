package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/lxn/walk"
)

func NewFooModel() *FooModel {
	m := new(FooModel)
	//m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *FooModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *FooModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Title
	case 1:
		if item.Tiempo.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			return item.Tiempo.Format("Today 15:04")
		}
		if item.Tiempo.Format("2006-01-02") == time.Now().AddDate(0, 0, -1).Format("2006-01-02") {
			return item.Tiempo.Format("Yesterday 15:04")
		}
		if item.Tiempo.After(time.Now().AddDate(0, 0, -7)) {
			return item.Tiempo.Format("Monday 15:04")
		}
		return item.Tiempo.Format("2006-01-02 15:04")
	case 2:
		return item.PP.PP.Total

	}
	panic("unexpected col")
}

func (m *FooModel) ResetRows() {
	foos := getReplays()
	m.items = make([]*Foo, len(foos))

	for i, foo := range foos {
		foo.Index = i
		m.items[i] = foo
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

// Called by the TableView to sort the model.
func (m *FooModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.Title < b.Title)

		case 1:
			return c(a.Tiempo.UnixNano() < b.Tiempo.UnixNano())
		case 2:
			return c(a.PP.PP.Total < b.PP.PP.Total)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func getImage(id int) walk.Image {
	im1, err := walk.NewImageFromFile(fmt.Sprintf(filepath.Join(osuFolder, "Data/bt/%dl.jpg"), id))
	if err != nil {
		return nil
	}
	return im1
}
