package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func getMainWindow(model *FooModel, tv *walk.TableView, imv *walk.ImageView, panelPP *PPanel) MainWindow {
	return MainWindow{
		Title:  "PIPO",
		Size:   Size{950, 500},
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			TableView{
				AssignTo: &tv,
				Columns: []TableViewColumn{
					{Title: "Title", Width: 350},
					{Title: "Date", Width: 105},
					{Title: "PP"},
				},
				Model:   model,
				MinSize: Size{600, 120},
				MaxSize: Size{600, 500},
				OnCurrentIndexChanged: func() {
					updateInfo(tv.CurrentIndex(), model, imv)
				},
			},
			Composite{
				AssignTo: &panel,
				Layout:   VBox{MarginsZero: true},
				MinSize:  Size{160, 0},
				MaxSize:  Size{400, 500},

				Children: []Widget{
					Composite{
						Layout:  HBox{MarginsZero: true},
						MaxSize: Size{400, 120},
						MinSize: Size{0, 120},
						Children: []Widget{
							ImageView{
								AssignTo: &imv,
								Image:    getImage(596704),
								MaxSize:  Size{160, 120},
								MinSize:  Size{160, 120},
							},
							Composite{
								Layout:  Grid{Columns: 2},
								MaxSize: Size{180, 120},
								Font:    Font{PointSize: 10},
								Children: []Widget{
									Label{
										Text: "Mods: ",
									},
									Label{
										AssignTo: &panelPP.Mods,
										Text:     "HDDT",
									},
									Label{
										Text: "Combo: ",
									},
									Label{
										AssignTo: &panelPP.Combo,
										Text:     "244/453x",
									},
									Label{
										Text: "Score: ",
									},
									Label{
										AssignTo: &panelPP.Score,
										Text:     "18.821.531",
									},
									Label{
										Text: "Accuracy: ",
									},
									Label{
										AssignTo: &panelPP.Acc,
										Text:     "99.33%",
									},
									Label{
										Text: "Rank: ",
										Font: Font{PointSize: 10},
									},
									Label{
										AssignTo:  &panelPP.Rank,
										Text:      "A",
										Font:      Font{PointSize: 12},
										TextColor: walk.RGB(0, 191, 0),
									},
								},
							},
						},
					},

					VSeparator{},
					Composite{
						Layout: HBox{MarginsZero: true},
						Font:   Font{PointSize: 11},
						Children: []Widget{
							Composite{
								Layout: Grid{Columns: 2},
								Font:   Font{PointSize: 10},
								Children: []Widget{
									Label{
										Text: "300s: ",
									},
									Label{
										AssignTo:  &panelPP.N300,
										Text:      "0",
										TextColor: walk.RGB(0, 0, 255),
									},
									Label{
										Text: "100s: ",
									},
									Label{
										AssignTo:  &panelPP.N100,
										Text:      "0",
										TextColor: walk.RGB(0, 191, 0),
									},
									Label{
										Text: "50s: ",
									},
									Label{
										AssignTo:  &panelPP.N50,
										Text:      "0",
										TextColor: walk.RGB(255, 144, 0),
									},
									Label{
										Text: "Misses: ",
									},
									Label{
										AssignTo:  &panelPP.Misses,
										Text:      "0",
										TextColor: walk.RGB(255, 0, 0),
									},
								},
							},
							GroupBox{
								Layout: Grid{Columns: 2},
								Font:   Font{PointSize: 9},
								Title:  "Difficulty",
								Children: []Widget{
									Label{
										Text: "Aim stars: ",
										Font: Font{PointSize: 11},
									},
									Label{
										AssignTo: &panelPP.AimStars,
										Text:     "0.00",
										Font:     Font{PointSize: 11},
									},
									Label{
										Text: "Speed stars: ",
										Font: Font{PointSize: 11},
									},
									Label{
										AssignTo: &panelPP.SpeedStars,
										Text:     "0.00",
										Font:     Font{PointSize: 11},
									},
									Label{
										Text: "Stars: ",
										Font: Font{PointSize: 11, Bold: true},
									},
									Label{
										AssignTo: &panelPP.Stars,
										Text:     "0.00",
										Font:     Font{PointSize: 11, Bold: true},
									},
								},
							},
							Composite{
								Layout: Grid{Columns: 2},
								Font:   Font{PointSize: 10},
								Children: []Widget{
									Label{
										Text: "AR: ",
									},
									Label{
										AssignTo: &panelPP.AR,
										Text:     "0",
									},
									Label{
										Text: "OD: ",
									},
									Label{
										AssignTo: &panelPP.OD,
										Text:     "0",
									},
									Label{
										Text: "CS: ",
									},
									Label{
										AssignTo: &panelPP.CS,
										Text:     "0",
									},
									Label{
										Text: "HP: ",
									},
									Label{
										AssignTo: &panelPP.HP,
										Text:     "0",
									},
								},
							},
						},
					},
					GroupBox{
						Layout: HBox{MarginsZero: true},
						Font:   Font{PointSize: 11},
						Children: []Widget{
							Composite{
								Layout: Grid{Columns: 1},
								Children: []Widget{
									Label{
										AssignTo: &panelPP.TotalPP,
										Text:     "0.0 pp",
										Font:     Font{PointSize: 16, Bold: true},
										RowSpan:  2,
									},
								},
							},
							Composite{
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{
										Text: "Acc PP: ",
									},
									Label{
										AssignTo: &panelPP.AccPP,
										Text:     "0.0 pp",
									},
									Label{
										Text: "Speed PP: ",
									},
									Label{
										AssignTo: &panelPP.SpeedPP,
										Text:     "0.0 pp",
									},
									Label{
										Text: "Aim PP: ",
									},
									Label{
										AssignTo: &panelPP.AimPP,
										Text:     "0.0 pp",
									},
								},
							},
						},
					},
					VSeparator{},
					GroupBox{
						Title:  "FC with same mods",
						Layout: Grid{Columns: 6},
						Font:   Font{PointSize: 10},
						Children: []Widget{
							Label{
								Text: "Accuracy:",
							},
							Label{
								Text: "95%",
							},
							Label{
								Text: "98%",
							},
							Label{
								Text: "99%",
							},
							Label{
								Text: "99.5%",
							},
							Label{
								Text: "100%",
							},
							Label{
								Text: "Total PP:",
							},
							Label{
								AssignTo: &panelPP.P95,
								Text:     "0.0", //95%
							},
							Label{
								AssignTo: &panelPP.P98,
								Text:     "0.0", // 98%
							},
							Label{
								AssignTo: &panelPP.P99,
								Text:     "0.0", // 99%
							},
							Label{
								AssignTo: &panelPP.P99p5,
								Text:     "0.0", // 99.50%
							},
							Label{
								AssignTo: &panelPP.P100,
								Text:     "0.0", // 100%
							},
						},
					},
				},
			},
		},
	}
}