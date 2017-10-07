package main

import (
	"fmt"
	"strconv"

	oppai "github.com/flesnuk/oppai5"
	"github.com/lxn/walk"
)

func updateInfo(i int, m *FooModel, im *walk.ImageView) {
	im.Synchronize(func() {
		old := im.Image()
		new := getImage(m.items[i].Foto)
		im.SetImage(new)
		if old != nil {
			old.Dispose()
		}
	})
	panel.Synchronize(func() {
		panelPP.Mods.SetText(oppai.ModsStr(int(m.items[i].Info.Mods)))
		panelPP.Combo.SetText(fmt.Sprintf("%d/%dx", m.items[i].Info.Combo, m.items[i].PP.Diff.Beatmap.MaxCombo))
		panelPP.Score.SetText(formatScore(int(m.items[i].Info.Score)))
		panelPP.N300.SetText(strconv.Itoa(int(m.items[i].Info.N300)))
		panelPP.N100.SetText(strconv.Itoa(int(m.items[i].Info.N100)))
		panelPP.N50.SetText(strconv.Itoa(int(m.items[i].Info.N50)))
		panelPP.Misses.SetText(strconv.Itoa(int(m.items[i].Info.Misses)))

		panelPP.AimStars.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Diff.Aim))
		panelPP.SpeedStars.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Diff.Speed))
		panelPP.Stars.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Diff.Total))

		panelPP.TotalPP.SetText(fmt.Sprintf("%.2f pp", m.items[i].PP.PP.Total))
		panelPP.AimPP.SetText(fmt.Sprintf("%.2f pp", m.items[i].PP.PP.Aim))
		panelPP.AccPP.SetText(fmt.Sprintf("%.2f pp", m.items[i].PP.PP.Acc))
		panelPP.SpeedPP.SetText(fmt.Sprintf("%.2f pp", m.items[i].PP.PP.Speed))

		panelPP.AR.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Stats.AR))
		panelPP.OD.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Stats.OD))
		panelPP.CS.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Stats.CS))
		panelPP.HP.SetText(fmt.Sprintf("%.2f", m.items[i].PP.Stats.HP))

		panelPP.P95.SetText(fmt.Sprintf("%.2f", m.items[i].PP.StepPP.P95))
		panelPP.P98.SetText(fmt.Sprintf("%.2f", m.items[i].PP.StepPP.P98))
		panelPP.P99.SetText(fmt.Sprintf("%.2f", m.items[i].PP.StepPP.P99))
		panelPP.P99p5.SetText(fmt.Sprintf("%.2f", m.items[i].PP.StepPP.P99p5))
		panelPP.P100.SetText(fmt.Sprintf("%.2f", m.items[i].PP.StepPP.P100))
		acc := (&oppai.Accuracy{int(m.items[i].Info.N300), int(m.items[i].Info.N100),
			int(m.items[i].Info.N50), int(m.items[i].Info.Misses)}).Value()
		panelPP.Acc.SetText(fmt.Sprintf("%.2f %%", acc*100.0))
		rankText := grade(m.items[i].Info)
		panelPP.Rank.SetText(rankText)
		panelPP.Rank.SetTextColor(gradeColor(rankText))
	})
}
