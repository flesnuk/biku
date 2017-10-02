package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/flesnuk/biku/osuhm"
	"github.com/flesnuk/osu-tools/osu"

	"github.com/lxn/walk"
)

const cachefile = "cache.gob"

var osuDirectory = "D:/osu!"

var hm *OsuHM

var panel *walk.Composite
var panelPP *PPanel

type lbl = *walk.Label

func getReplays() []*Foo {
	list, _ := ReadDirByTime(filepath.Join(osuDirectory, "Data/r"))
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
	hm = Load(".")
	var tv *walk.TableView = new(walk.TableView)
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
