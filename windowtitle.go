package main

import (
	"github.com/mitchellh/go-ps"
)

func isOsuOpen() bool {
	processName := "osu!.exe"
	processes, err := ps.Processes()

	if err != nil {
		return false
	}

	for _, p := range processes {
		if p.Executable() == processName {
			return true
		}
	}
	return false
}
