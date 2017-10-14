package main

import (
	"time"

	oppai "github.com/flesnuk/oppai5"
	"github.com/flesnuk/osu-tools/osu"
	"github.com/lxn/walk"
)

type PPanel struct {
	Mods, Combo, Score, Acc, Rank  lbl
	N300, N100, N50, Misses        lbl
	AimStars, SpeedStars, Stars    lbl
	AR, OD, CS, HP                 lbl
	TotalPP, AccPP, SpeedPP, AimPP lbl
	P95, P98, P99, P99p5, P100     lbl
}

type Row struct {
	Index  int
	Title  string
	Foto   int
	Tiempo time.Time
	PP     oppai.PP
	Info   osu.Replay
}

type RowModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Row
}
