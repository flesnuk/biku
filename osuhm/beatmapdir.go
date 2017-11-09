package osuhm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
)

/*
	Created by https://github.com/Lubieerror provided in https://github.com/flesnuk/biku/issues/1

	Steps:
	1. Find Windows username
	2. Find osu!.%windows_filename%.cfg file
	3. Find "BeatmapDirectory" value (and "clear" it) or search if default folder exists
	4. Go to biku/osuhm/osuhm.go, to GetBeatmapPath func
	5. Change string "Songs" to value from 3'rd point

	If it's needed, there is no license for this file and code etc. You can also treat this like a WTFPL or smt
*/

func (osuhm *OsuHM) InitBeatmapDir() {
	BeatmapDir = osuhm.getBeatmapDirectory(getMachineUsername())
}

// GetBeatmapDirectory returns the "Songs" folder from BeatmapDirectory conf variable resided in .cfg file
func (osuhm *OsuHM) getBeatmapDirectory(username string) string {
	f, err := os.Open(path.Join(osuhm.OsuFolder, "osu!."+username+".cfg"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "BeatmapDirectory") {
			// Get the BM location from setting BeatmapDirectory and remove spaces !WARNING! I didn't
			// tested how spaces inside path works at all (if Osu reads for eg. "D:\My Osu Songs\Songs",
			// if Trim function won't delete them, if it'll be working with quotes etc.), so pls test
			// before merge! :)
			return strings.TrimSpace(strings.Split(line, "=")[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// If theres no "Beatmap Directory" setting, try default "Songs" folder
	// Could also return "" and then/instead read error, or check if %OsuFolder%/Songs exists
	// or something... Just preferences.
	return "Songs"
}

func getMachineUsername() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get the username and get rid of computer name or smt (be careful don't use "user.Name"! It could be
	// same as username (eg."jankow") but it's common to give it some sort of full name eg. "Jan Kowalsky")
	// I'm sure you know that but better safe than sorry (and waste time debugging) + you could have same name and username
	return strings.Split(user.Username, "\\")[1]
}
