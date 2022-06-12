//go:build windows

package tools

import (
	"flag"
	"fmt"
)

var (
	FlDebug = flag.Bool("debug", false, "Enable debug mode")
)

// IsError returns true if an error occurs and prints error message if -debug flag is used
func IsError(e error, msg ...string) bool {
	if e != nil {
		Debug(msg, e)
		return true
	}
	return false
}

// Debug prints out messages if -debug flag is used
func Debug(a ...any) {
	if *FlDebug {
		fmt.Println(a...)
	}
}

// Console shows/hides a command line window if set to true/false
func Console(show bool) {
	var getWin = Kernel32.MustFindProc("GetConsoleWindow")
	var showWin = User32.MustFindProc("ShowWindow")
	hwnd, _, err := getWin.Call()
	IsError(err, "Console() - getWin.Call()")
	if hwnd == 0 {
		return
	}
	if show {
		var swRestore uintptr = 9
		_, _, err := showWin.Call(hwnd, swRestore)
		IsError(err, "Console() - showWin.Call(hwnd, swRestore)")

	} else {
		var swHide uintptr = 0
		_, _, err := showWin.Call(hwnd, swHide)
		IsError(err, "Console() - showWin.Call(hwnd, swHide)")
	}
}
