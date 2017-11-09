package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func saveLogIfPanic() {
	if x := recover(); x != nil {
		// recovering from a panic; x contains whatever was passed to panic()
		file, err := os.Create("biku-log-" + time.Now().Format("20060102-15.04.05") + ".txt")
		if err != nil {
			fmt.Println(err)
		}
		file.WriteString(fmt.Sprint(x, "\r\n\r\n"))
		file.WriteString(strings.Replace(string(debug.Stack()), "\n", "\r\n", -1))
		panic(x)
	}
}
