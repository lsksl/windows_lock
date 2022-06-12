//go:build windows

package tools

import (
	"syscall"
	"time"
	"unsafe"
)

//const (
//	DESKTOP_SWITCHDESKTOP = 0x0100 // The access to the desktop
//)

var (
	User32           = syscall.MustLoadDLL("user32.dll")
	Kernel32         = syscall.MustLoadDLL("Kernel32.dll")
	getLastInputInfo = User32.MustFindProc("GetLastInputInfo")
	getTickCount     = Kernel32.MustFindProc("GetTickCount")
	lockWorkstation  = User32.MustFindProc("LockWorkStation")
	lastInputInfo    struct {
		cbSize uint32
		dwTime uint32
	}
)

func LockWindows() {
	_, _, err := lockWorkstation.Call()
	IsError(err, "LockWindows()")
}

func IdleTime() time.Duration {
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	currentTickCount, _, _ := getTickCount.Call()
	r1, _, err := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r1 == 0 {
		IsError(err, "IdleTime()")
	}
	return time.Duration(uint32(currentTickCount)-lastInputInfo.dwTime) * time.Millisecond
}
