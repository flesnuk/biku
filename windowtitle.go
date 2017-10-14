package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// credits to @txanatan https://github.com/txanatan/ppwatch/blob/master/main.go

const getWindowTitlePowershellExpression = `gps |
? {$_.mainwindowtitle -like "*%s*"} |
? {$_.processname -like "%s"} |
select mainwindowtitle`

func isOsuOpen() bool {
	PartialWindowTitle := "osu!"
	ProcessName := "osu!"
	expr := fmt.Sprintf(getWindowTitlePowershellExpression, PartialWindowTitle, ProcessName)
	cmd := exec.Command("powershell", "-command", expr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	// Strip all leading/trailing whitespace, separate into lines, delete empty lines
	var outlines []string
	for _, s := range strings.Split(strings.Trim(string(out), " "), "\r\n") {
		if s != "" {
			outlines = append(outlines, s)
		}
	}

	// Check that we have content in the array, if not the window doesn't exist
	if len(outlines) < 1 {
		return false
	}
	return true
}
