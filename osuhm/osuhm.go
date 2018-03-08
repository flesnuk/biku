package osuhm

import (
	"strconv"
	"github.com/karrick/godirwalk"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/flesnuk/osu-tools/osu"
	"github.com/flesnuk/osu-tools/osudb"
)

const cacheName = "biku-cache.gob"

var BeatmapDir = "Songs"

// OsuHM stores osu root path, the api key and hashmap of beatmaps
type OsuHM struct {
	OsuFolder string
	APIKey    string
	HM        map[string]osu.Beatmap
}

// New creates a new OsuHM
// it opens osu!.db file, if it was opened before use Load instead
func NewOsuHM(osuFolder string) *OsuHM {
	osudbFile, err := os.Open(filepath.Join(osuFolder, "osu!.db"))
	defer osudbFile.Close()
	if err != nil {
		return nil
	}
	hm := osudb.GetBeatmaps(osudbFile)
	return &OsuHM{
		OsuFolder: osuFolder,
		HM:        hm,
	}
}

// Load loads a previously cached OsuHM
// dirPath is the folder where the cache file is located
func Load(dirPath string) *OsuHM {
	ret := &OsuHM{}
	err := load(path.Join(dirPath, cacheName), ret)
	if err != nil {
		return nil
	}
	return ret
}

// SaveCache stores an OsuHM cache file in dest folder
func (osuhm *OsuHM) SaveCache(dest string) {
	err := save(path.Join(dest, cacheName), osuhm)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetBeatmap returns a osu! Beatmap from its hash
func (osuhm *OsuHM) GetBeatmap(beatmapHash string) *osu.Beatmap {
	bm, ok := osuhm.HM[beatmapHash]

	if !ok {
		if osuhm.APIKey == "" {
			return nil
		}
		bm = apiGetBeatmap(osuhm.APIKey, beatmapHash)
		if bm.ID == 0 {
			return nil
		}
		osuhm.HM[beatmapHash] = bm
	}

	return &bm
}

// StartNotifier runs a notifier that watch for new replays and
// sends them into the replayChan channel
func (osuhm *OsuHM) StartNotifier(replayChan chan osu.Replay) {
	go RunNotifier(osuhm.OsuFolder, replayChan)
}

// GetBeatmapPath returns the full path of the specified beatmap
func (osuhm *OsuHM) GetBeatmapPath(beatmap *osu.Beatmap) string {
	var beatmapDir string
	if BeatmapDir == "Songs" {
		beatmapDir = filepath.Join(osuhm.OsuFolder, BeatmapDir)
	} else {
		beatmapDir = filepath.ToSlash(BeatmapDir)
	}
	
	children, err := godirwalk.ReadDirnames(beatmapDir, nil)
	if err != nil {
		return ""
	}

	beatmapPath := ""
	for _, child := range children {
		if strings.HasPrefix(child, strconv.Itoa(int(beatmap.ID))) {
			beatmapPath = filepath.Join(beatmapDir, child)
		}
	}
	
	return filepath.Join(beatmapPath,
		beatmap.Filename)

}
