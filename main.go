package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"time"
	"windows_lock/icon"
	"windows_lock/tools"
)

var (
	lockOptions tools.LockTimeOptions
	settings    tools.Settings
	err         error
)

func main() {
	lockOptions, settings, err = tools.ReadConfig()

	_ = tools.IsError(err)
	fmt.Println(settings.LockTimer)
	fmt.Println(lockOptions)

	systray.Run(onReady, onExit)
}

func onExit() {
	systray.Quit()
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTooltip("Lantern")
	mLockTime := systray.AddMenuItem("Lock screen after...", "")
	mLock := systray.AddMenuItem("Lock screen now", "")
	mQuit := systray.AddMenuItem("Quit", "")

	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T1), "", true)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T2), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T3), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T4), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T5), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T6), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T7), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T8), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T9), "", false)
	mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T10), "", false)

	t := time.NewTicker(1 * time.Second)

	go func() {
		for range t.C {
			if tools.IdleTime() < 1000000000 {
				fmt.Printf("***************Countdown reset at %s ***************\n", time.Now())
			}
			fmt.Println(tools.IdleTime())
		}
	}()

	for {
		select {
		case <-mQuit.ClickedCh:
			onExit()
		case <-mLock.ClickedCh:
			lockScreen()
		}
	}
}

func lockScreen() {
	robotgo.KeySleep = 100
	robotgo.KeyToggle("cmd")
	robotgo.KeyTap("l")
}
