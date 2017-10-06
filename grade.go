package main

import (
	"github.com/flesnuk/osu-tools/osu"
	"github.com/lxn/walk"
)

// https://github.com/chudooder/osutools/blob/master/dbparse.py#L249-L264
func grade(r osu.Replay) string {
	numNotes := r.Misses + r.N300 + r.N100 + r.N50
	wScore := r.N300 + r.N100*2.0/6.0 + r.N50*1.0/6.0
	acc := wScore / numNotes
	switch {
	case acc == 1.0:
		return "SS"
	case float32(r.N300)/float32(numNotes) >= 0.9 &&
		float32(r.N50)/float32(numNotes) <= 0.1 &&
		r.Misses == 0:
		return "S"
	case float32(r.N300)/float32(numNotes) >= 0.8 &&
		float32(r.N300)/float32(numNotes) >= 0.9:
		return "A"
	case float32(r.N300)/float32(numNotes) >= 0.7 && r.Misses == 0 ||
		float32(r.N300)/float32(numNotes) >= 0.8:
		return "B"
	case float32(r.N300)/float32(numNotes) >= 0.6:
		return "C"
	}

	return "D"
}

func gradeColor(grade string) walk.Color {
	rgb := walk.RGB
	switch grade {
	case "SS":
		return rgb(255, 208, 20)
	case "S":
		return rgb(255, 208, 20)
	case "A":
		return rgb(57, 216, 26)
	case "B":
		return rgb(15, 74, 224)
	case "C":
		return rgb(249, 19, 246)
	case "D":
		return rgb(193, 21, 41)
	default:
		return rgb(0, 0, 0)
	}
}
