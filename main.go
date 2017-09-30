package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	. "github.com/flesnuk/biku/osuhm"

	"github.com/lxn/walk"
)

const cachefile = "cache.gob"

var osuDirectory = "D:/osu!"

var hm *OsuHM

var panel *walk.Composite
var panelPP *PPanel

type lbl = *walk.Label

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by modtime.
func ReadDirByTime(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ModTime().UnixNano() > list[j].ModTime().UnixNano() })
	return list, nil
}

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

	panelPP = new(PPanel)

	var imv *walk.ImageView = new(walk.ImageView)

	getMainWindow(model, tv, imv, panelPP).Run()

}
