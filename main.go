package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	. "github.com/flesnuk/biku/osuhm"
	oppai "github.com/flesnuk/oppai5"
	"github.com/flesnuk/osu-tools/osu"

	"github.com/lxn/walk"
)

const cachefile = "cache.gob"

var osuFolder = ""
var lastago = -1

var hm *OsuHM
var mw *walk.MainWindow

var panel *walk.Composite
var panelPP *PPanel

type lbl = *walk.Label

func calcPP(osuFile *os.File, replay osu.Replay, row *Row) {
	row.PP = oppai.PPInfo(oppai.Parse(osuFile), &oppai.Parameters{
		replay.N300,
		replay.N100,
		replay.N50,
		replay.Misses,
		replay.Combo,
		replay.Mods,
	})
	osuFile.Close()
}

func main() {
	var err error
	defer saveLogIfPanic()

	tv := new(walk.TableView)

	hm = Load(".")
	if hm == nil {
		if isOsuOpen() {
			walk.MsgBox(nil, "osu!db", "Please, close osu! before starting this app for the first time",
				walk.MsgBoxIconExclamation)
			return
		}
		var ok bool
		osuFolder, ok = checkAll()
		for !ok {
			fd := new(walk.FileDialog)
			fd.Title = "Select your osu! folder"
			accepted, err := fd.ShowBrowseFolder(nil)
			if !accepted || err != nil {
				os.Exit(0)
			}
			osuFolder = fd.FilePath
			_, ok = checkAll()
		}
		hm, err = NewOsuHM(osuFolder)
		if err != nil {
			walk.MsgBox(nil, "osu!db", err.Error(), walk.MsgBoxIconError)
		}
		osuapi := &OsuAPIKey{""}
		if _, err := getDialog(osuapi).Run(nil); err != nil {
			fmt.Println(err)
		}
		hm.APIKey = osuapi.OsuAPI
		hm.SaveCache(".")
	}

	err = hm.InitBeatmapDir()
	if err != nil {
		walk.MsgBox(nil, "InitBeatmapDir", err.Error(), walk.MsgBoxIconError)
	}

	model := NewRowModel()
	tv.Synchronize(func() {
		model.Sort(1, walk.SortDescending)
	})

	replayChan := make(chan osu.Replay)
	hm.StartNotifier(replayChan)

	go func() {
		for replay := range replayChan {
			bm := hm.GetBeatmap(replay.BeatmapHash)
			if bm == nil {
				continue
			}

			osuFile, err := os.Open(hm.GetBeatmapPath(bm))
			if err != nil {
				continue
			}

			row := createRow(osuFile, &replay, bm)
			model.items = append(model.items, row)
			calcPP(osuFile, replay, row)

			model.ResetRows()
			tv.Synchronize(func() {
				model.Sort(1, walk.SortDescending)
			})
		}

	}()

	panelPP = new(PPanel)

	imv := new(walk.ImageView)
	window := getMainWindow(model, tv, imv, panelPP)
	window.Create()
	data, err := Asset("icon/biku.png")
	if err != nil || len(data) == 0 {
		fmt.Println("Asset was not found.")
	}
	im, err := png.Decode(bytes.NewReader(data))
	ic, err := walk.NewIconFromImage(im)
	mw.SetIcon(ic)

	mw.Run()
	hm.SaveCache(".")
}
