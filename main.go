package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/flesnuk/biku/osuhm"
	"github.com/flesnuk/osu-tools/osu"

	"github.com/lxn/walk"
)

const cachefile = "cache.gob"

var osuFolder = "D:/osu"

var hm *OsuHM

var panel *walk.Composite
var panelPP *PPanel

type lbl = *walk.Label

func getReplays() []*Foo {
	list, _ := ReadDirByTime(filepath.Join(osuFolder, "Data/r"))
	ret := make([]*Foo, 0, 5)
	for _, x := range list {
		if !strings.HasSuffix(x.Name(), "osr") {
			continue
		}
		if x.ModTime().After(time.Now().AddDate(0, 0, 0)) {
			continue
		}
		if x.ModTime().Before(time.Now().AddDate(0, 0, -7)) {
			break
		}

		replay := getReplay(x)
		if replay == nil {
			continue
		}
		bm := hm.GetBeatmap(replay.BeatmapHash)
		if bm == nil {
			continue
		}

		osuFile, err := os.Open(hm.GetBeatmapPath(bm))
		if err != nil {
			continue
		}

		ret = append(ret, createFoo(osuFile, replay, bm))
		osuFile.Close()

	}
	return ret

}

func main() {
	var tv *walk.TableView = new(walk.TableView)

	fd := new(walk.FileDialog)

	_, ok := checkAll()
	for !ok {
		accepted, err := fd.ShowBrowseFolder(nil)
		if !accepted || err != nil {
			os.Exit(0)
		}
		osuFolder = fd.FilePath
		_, ok = checkAll()
		fmt.Println(ok)
	}

	hm = Load(".")
	if hm == nil {
		hm = NewOsuHM(osuFolder)
		hm.SaveCache(".")
	}

	model := NewFooModel()
	tv.Synchronize(func() {
		model.Sort(1, walk.SortDescending)
	})

	x := make(chan osu.Replay)

	hm.StartNotifier(x)

	go func() {
		for replay := range x {
			replay.ModTime = time.Now()
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

	var imv *walk.ImageView = new(walk.ImageView)

	getMainWindow(model, tv, imv, panelPP).Run()

}
