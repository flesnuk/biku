package osuhm

import (
	"fmt"
	"strings"

	"github.com/flesnuk/osu-tools/osu"
	"gopkg.in/thehowl/go-osuapi.v1"
)

func apiGetBeatmap(apiKey, beatmapHash string) *osu.Beatmap {
	c := osuapi.NewClient(apiKey)
	bms, err := c.GetBeatmaps(osuapi.GetBeatmapsOpts{
		BeatmapHash: beatmapHash,
	})

	if err != nil || len(bms) <= 0 {
		return nil
	}

	bm := bms[0]
	fname := fmt.Sprintf("%s - %s (%s) [%s].osu",
		bm.Artist, bm.Title, bm.Creator, bm.DiffName)
	return &osu.Beatmap{
		ID:       uint32(bm.BeatmapSetID),
		Filename: strings.Replace(fname, "/", "", -1),
	}

}
