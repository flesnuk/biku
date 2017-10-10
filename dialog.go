package main

import (
	"os/exec"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type OsuAPIKey struct {
	OsuAPI string
}

func getDialog(osuapi *OsuAPIKey) Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	return Dialog{
		Title:         "osu! API key",
		AssignTo:      &dlg,
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		MinSize:       Size{320, 180},
		DataBinder: DataBinder{
			AssignTo:       &db,
			DataSource:     osuapi,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			Label{
				Text: "Put your osu! API Key here. It will be used ",
			},
			Label{
				Text: "when you download new maps and the cache isn't up to date",
			},
			LinkLabel{
				Text: `(if you don't have one, you can get one <a id="this" href="http://osu.ppy.sh/p/api">here</a>)`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					exec.Command("rundll32", "url.dll,FileProtocolHandler", link.URL()).Start()
				},
			},
			VSpacer{},
			LineEdit{
				Text: Bind("OsuAPI"),
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								return
							}
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
