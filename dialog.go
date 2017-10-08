package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func getDialog() Dialog {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		Title:         "osu! API key",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			LineEdit{
				Text: "Insert your osu! API key here",
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}
}
