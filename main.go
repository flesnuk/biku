package main

import (
	"fmt"
	"os"
	"path"
	"time"

	. "github.com/flesnuk/biku/osuhm"
	"github.com/flesnuk/osu-tools/osr"
	"github.com/flesnuk/osu-tools/osu"

	"github.com/lxn/walk"
)

const cachefile = "cache.gob"

var osuFolder = ""
var lastago = -1

var hm *OsuHM

var panel *walk.Composite
var panelPP *PPanel

type lbl = *walk.Label

func getReplays() []*Foo {
	ff, err := os.Open(path.Join(osuFolder, "scores.db"))
	defer ff.Close()
	if err != nil {
		fmt.Println("FAIL")
	}
	list := osr.ReadScoreDB(ff)
	//list, _ := ReadDirByTime(filepath.Join(osuFolder, "Data/r"))
	ret := make([]*Foo, 0, 5)
	for _, replay := range list {
		if replay.ModTime.After(time.Now().AddDate(0, 0, 0)) {
			continue
		}
		if replay.ModTime.Before(time.Now().AddDate(0, 0, lastago)) {
			break
		}

		bm := hm.GetBeatmap(replay.BeatmapHash)
		if bm == nil {
			continue
		}

		osuFile, err := os.Open(hm.GetBeatmapPath(bm))
		if err != nil {
			continue
		}
		ret = append(ret, createFoo(osuFile, &replay, bm))
		osuFile.Close()

	}
	return ret

}

func main() {
	tv := new(walk.TableView)

	// if cmd, err := getDialog().Run(nil); err != nil {
	// 	fmt.Println(err)
	// } else if cmd == walk.DlgCmdOK {
	// 	fmt.Println("OK")
	// } else if cmd == walk.DlgCmdCancel {
	// 	fmt.Println("Cancel")
	// }

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
		hm = NewOsuHM(osuFolder)
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

			model.items = append(model.items, createFoo(osuFile, &replay, bm))
			osuFile.Close()

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
