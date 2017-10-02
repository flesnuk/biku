package main

import (
	"os"
	"path/filepath"
	"sort"

	oppai "github.com/flesnuk/oppai5"
	"github.com/flesnuk/osu-tools/osr"
	"github.com/flesnuk/osu-tools/osu"
)

func getReplay(x os.FileInfo) *osu.Replay {
	f, err := os.Open(filepath.Join(filepath.Join(hm.OsuFolder, "Data/r"), x.Name()))
	if err != nil {
		return nil
	}
	replay := osr.NewReplay(f)
	f.Close()
	if replay.GameMode != 0 {
		return nil
	}
	replay.ModTime = x.ModTime()
	return &replay
}

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

func createFoo(osuFile *os.File, replay *osu.Replay, bm *osu.Beatmap) *Foo {
	pp := oppai.PPInfo(oppai.Parse(osuFile), &oppai.Parameters{
		replay.N300,
		replay.N100,
		replay.N50,
		replay.Misses,
		replay.Combo,
		replay.Mods,
	})

	return &Foo{
		Title:  bm.Filename,
		Foto:   int(bm.ID),
		Tiempo: replay.ModTime,
		PP:     pp,
		Info:   *replay,
	}
}
