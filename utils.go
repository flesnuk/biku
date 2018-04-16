package main

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
	"sort"
	"strconv"

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
	replay.Path = path.Join(hm.OsuFolder, "Data/r", x.Name())
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

func createRow(osuFile *os.File, replay *osu.Replay, bm *osu.Beatmap) *Row {
	return &Row{
		Title:  bm.Filename,
		Foto:   int(bm.ID),
		Tiempo: replay.ModTime,
		Info:   *replay,
	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func check(osuFolder string) bool {
	if exists(osuFolder) &&
		exists(path.Join(osuFolder, "osu!.exe")) &&
		exists(path.Join(osuFolder, "osu!.db")) &&
		exists(path.Join(osuFolder, "Data/r")) {
		return true
	}
	return false
}

func checkAll() (string, bool) {
	switch {
	case check(osuFolder):
		return osuFolder, true
	case check(path.Join("C:", "Program Files", "osu!")):
		return path.Join("C:", "Program Files", "osu!"), true
	case check(path.Join("C:", "Program Files (x86)", "osu!")):
		return path.Join("C:", "Program Files (x86)", "osu!"), true
	case check("."):
		return ".", true
	default:
		if usr, err := user.Current(); err == nil {
			f := path.Join(usr.HomeDir, "AppData", "Local", "osu!")
			return f, check(f)
		}
		return "", false
	}
}

func formatScore(n int) string {
	fstr := ""
	for n >= 1000 {
		if x := n % 1000; x != 0 {
			switch {
			case x < 10:
				fstr = ".00" + strconv.Itoa(x) + fstr
			case x < 100:
				fstr = ".0" + strconv.Itoa(x) + fstr
			default:
				fstr = "." + strconv.Itoa(x) + fstr
			}

		} else {
			fstr = ".000" + fstr
		}
		n /= 1000
	}
	return strconv.Itoa(n) + fstr
}
