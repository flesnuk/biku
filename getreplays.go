package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flesnuk/osu-tools/osu"

	"github.com/flesnuk/osu-tools/osr"
)

func getReplaysFromDB() []*Row {
	ff, err := os.Open(path.Join(hm.OsuFolder, "scores.db"))
	defer ff.Close()
	if err != nil {
		fmt.Println("Couldn't open scores.db")
	}
	list := osr.ReadScoreDB(ff)
	ret := make([]*Row, 0, 5)
	for _, replay := range list {
		if replay.ModTime.After(time.Now().AddDate(0, 0, 0)) {
			continue
		}
		if replay.ModTime.Before(time.Now().AddDate(0, 0, lastago)) {
			break
		}
		replayPath := replayPath(&replay)
		if _, err := os.Stat(replayPath); err == nil {
			replay.Path = replayPath
		}

		bm := hm.GetBeatmap(replay.BeatmapHash)
		if bm == nil {
			continue
		}

		osuFile, err := os.Open(hm.GetBeatmapPath(bm))
		if err != nil {
			continue
		}

		row := createRow(osuFile, &replay, bm)
		ret = append(ret, row)
		calcPP(osuFile, replay, row)

	}
	return ret

}

func getReplays() []*Row {
	list, _ := ReadDirByTime(filepath.Join(hm.OsuFolder, "Data/r"))
	ret := make([]*Row, 0, 5)
	for _, x := range list {
		if !strings.HasSuffix(x.Name(), "osr") {
			continue
		}
		if x.ModTime().After(time.Now().AddDate(0, 0, 0)) {
			continue
		}
		if x.ModTime().Before(time.Now().AddDate(0, 0, lastago)) {
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
		
		row := createRow(osuFile, replay, bm)
		ret = append(ret, row)
		calcPP(osuFile, *replay, row)

	}
	return ret

}

func replayPath(replay *osu.Replay) string {
	filename := replay.BeatmapHash + "-" + strconv.FormatInt(int64(replay.TimeStamp)-504911232000000000, 10) + ".osr"
	return path.Join(hm.OsuFolder, "Data", "r", filename)
}
