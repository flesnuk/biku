package main

import (
	"fmt"
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

func calcPP(osuFile *os.File, replay osu.Replay, foo *Foo) {
	foo.PP = oppai.PPInfo(oppai.Parse(osuFile), &oppai.Parameters{
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
	tv := new(walk.TableView)

	hm = Load(".")
	if hm == nil {
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
		hmaux := NewOsuHM(osuFolder)
		if hmaux == nil {
			walk.MsgBox(mw, "osu!db", "Please, close osu! before refreshing the cache",
				walk.MsgBoxIconExclamation)
			return
		}
		hm = hmaux
		osuapi := &OsuAPIKey{""}
		if _, err := getDialog(osuapi).Run(nil); err != nil {
			fmt.Println(err)
		}
		hm.APIKey = osuapi.OsuAPI
		hm.SaveCache(".")
	}

	osuFolder = hm.OsuFolder

	model := NewFooModel()
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

			foo := createFoo(osuFile, &replay, bm)
			model.items = append(model.items, foo)
			calcPP(osuFile, replay, foo)

			tv.Synchronize(func() {
				model.Sort(1, walk.SortDescending)
			})
		}

	}()

	panelPP = new(PPanel)

	imv := new(walk.ImageView)
	getMainWindow(model, tv, imv, panelPP).Run()
	hm.SaveCache(".")
}

//504911232000000000 y 1601
