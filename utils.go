package main

import (
	"os"
	"path/filepath"

	oppai "github.com/flesnuk/oppai5"
	"github.com/flesnuk/osu-tools/osr"
	"github.com/flesnuk/osu-tools/osu"
)

func getReplay(x os.FileInfo) *osu.Replay {
	osuDirectory := "D:/osu!"
	f, err := os.Open(filepath.Join(filepath.Join(osuDirectory, "Data/r"), x.Name()))
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
