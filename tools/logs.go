package tools

import (
	"flag"
	"fmt"
	"syscall"
)

var (
	FlDebug = flag.Bool("debug", false, "Enable debug mode")
)

// IsError returns true if an error occurs and prints error message to Debug
func IsError(e error) bool {
	if e != nil {
		Debug(e)
		return true
	}
	return false
}

// Debug prints out messages if -debug flag is used
func Debug(a ...any) {
	if *FlDebug {
		fmt.Println(a)
	}
}

// Console shows/hides a command line window if set to true/false
func Console(show bool) {
	var getWin = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	var showWin = syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	hwnd, _, _ := getWin.Call()
	if hwnd == 0 {
		return
	}
	if show {
		var swRestore uintptr = 9
		_, _, err := showWin.Call(hwnd, swRestore)
		IsError(err)

	} else {
		var swHide uintptr = 0
		_, _, err := showWin.Call(hwnd, swHide)
		IsError(err)
	}
}
