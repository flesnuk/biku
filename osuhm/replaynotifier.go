package osuhm

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/flesnuk/osu-tools/osr"

	"github.com/flesnuk/osu-tools/osu"
	"github.com/rjeczalik/notify"
)

// RunNotifier needs to be run concurrently for sending replays
// to replayChan when new replays are created in Data/r folder
func RunNotifier(osuFolder string, replayChan chan osu.Replay) {
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening on events within Data/r directory.
	// Dispatch each create events separately to c.
	if err := notify.Watch(path.Join(osuFolder, "Data/r"), c, notify.Create); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	for ei := range c {
		if !strings.HasSuffix(ei.Path(), "osr") {
			continue
		}

		time.Sleep(time.Millisecond * 100)
		var f *os.File
		var err error
		// Try opening the replay file until osu! closes it
		for attempts := 10; attempts > 0; attempts-- {
			f, err = os.Open(filepath.ToSlash(ei.Path()))
			if err == nil {
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
		replay, err := osr.NewReplay(f)
		f.Close()
		if err != nil {
			continue
		}
		replay.ModTime = time.Now()
		replay.Path = filepath.ToSlash(ei.Path())
		replayChan <- replay
	}
}
