//go:build windows && amd64

package main

import (
	"flag"
	"fmt"
	"github.com/getlantern/systray"
	"time"
	"windows_lock/icon"
	"windows_lock/tools"
)

var (
	lockOptions          tools.LockTimeOptions
	settings             tools.Settings
	err                  error
	defaultIdleReset     = 4
	lockTickerCount      = defaultIdleReset
	idleTicker           *time.Ticker
	lockTickerTime       time.Duration
	lockTicker           *time.Ticker
	resetCountTickerTime time.Duration
	resetCountTicker     *time.Ticker
)

func main() {
	flag.Parse()
	// Hide console if -debug flag is not used
	if !*tools.FlDebug {
		tools.Console(true)
	}

	// Read config file
	lockOptions, settings, err = tools.ReadConfig()
	_ = tools.IsError(err)
	tools.Debug("settings.LockTimer =", settings.LockTimer)
	tools.Debug("lockOptions =", lockOptions)

	systray.Run(onReady, onExit)
}

func onExit() {
	systray.Quit()
}

func onReady() {
	systray.SetIcon(icon.Data)
	mMsg := systray.AddMenuItem("", "")
	updateLockMessage(settings.LockTimer, mMsg)
	mLockTime := systray.AddMenuItem("Set screen lock to", "")
	mLock := systray.AddMenuItem("Lock screen now", "")
	mQuit := systray.AddMenuItem("Quit", "")

	mT1 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T1), "", lockOptions.T1 == settings.LockTimer)
	mT2 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T2), "", lockOptions.T2 == settings.LockTimer)
	mT3 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T3), "", lockOptions.T3 == settings.LockTimer)
	mT4 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T4), "", lockOptions.T4 == settings.LockTimer)
	mT5 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T5), "", lockOptions.T5 == settings.LockTimer)
	mT6 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T6), "", lockOptions.T6 == settings.LockTimer)
	mT7 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T7), "", lockOptions.T7 == settings.LockTimer)
	mT8 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T8), "", lockOptions.T8 == settings.LockTimer)
	mT9 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T9), "", lockOptions.T9 == settings.LockTimer)
	mT10 := mLockTime.AddSubMenuItemCheckbox(fmt.Sprintf("%d minutes", lockOptions.T10), "", lockOptions.T10 == settings.LockTimer)
	mT0 := mLockTime.AddSubMenuItemCheckbox("Don't lock", "", settings.LockTimer == 0)

	go startLockTicker()
	go startCountTicker()
	go startIdleTicker()

	for {
		select {
		case <-mQuit.ClickedCh:
			onExit()
		case <-mLock.ClickedCh:
			tools.LockWindows()
		case <-mT0.ClickedCh:
			tools.SetLockTimer(0, &settings)
			updateLockMessage(0, mMsg)
			mT0.Check()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT1.ClickedCh:
			tools.SetLockTimer(lockOptions.T1, &settings)
			updateLockMessage(lockOptions.T1, mMsg)
			mT0.Uncheck()
			mT1.Check()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT2.ClickedCh:
			tools.SetLockTimer(lockOptions.T2, &settings)
			updateLockMessage(lockOptions.T2, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Check()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT3.ClickedCh:
			tools.SetLockTimer(lockOptions.T3, &settings)
			updateLockMessage(lockOptions.T3, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Check()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT4.ClickedCh:
			tools.SetLockTimer(lockOptions.T4, &settings)
			updateLockMessage(lockOptions.T4, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Check()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT5.ClickedCh:
			tools.SetLockTimer(lockOptions.T5, &settings)
			updateLockMessage(lockOptions.T5, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Check()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT6.ClickedCh:
			tools.SetLockTimer(lockOptions.T6, &settings)
			updateLockMessage(lockOptions.T6, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Check()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT7.ClickedCh:
			tools.SetLockTimer(lockOptions.T7, &settings)
			updateLockMessage(lockOptions.T7, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Check()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT8.ClickedCh:
			tools.SetLockTimer(lockOptions.T8, &settings)
			updateLockMessage(lockOptions.T8, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Check()
			mT9.Uncheck()
			mT10.Uncheck()
		case <-mT9.ClickedCh:
			tools.SetLockTimer(lockOptions.T9, &settings)
			updateLockMessage(lockOptions.T9, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Check()
			mT10.Uncheck()
		case <-mT10.ClickedCh:
			tools.SetLockTimer(lockOptions.T10, &settings)
			updateLockMessage(lockOptions.T10, mMsg)
			mT0.Uncheck()
			mT1.Uncheck()
			mT2.Uncheck()
			mT3.Uncheck()
			mT4.Uncheck()
			mT5.Uncheck()
			mT6.Uncheck()
			mT7.Uncheck()
			mT8.Uncheck()
			mT9.Uncheck()
			mT10.Check()
		}
	}
}

func updateLockMessage(t uint16, m *systray.MenuItem) {
	msg := fmt.Sprintf("Screen will be locked after %d", t)
	if t == 0 {
		msg = fmt.Sprintf("Screen auto lock disabled")
	}
	systray.SetTooltip(msg)
	m.SetTitle(msg)
}

func startLockTicker() {
	lockTickerTime = checkLockTimer() * time.Minute
	lockTicker = time.NewTicker(lockTickerTime)
	for {
		tools.Debug("[startLockTicker()] lockTickerCount =", lockTickerCount)
		tools.Debug("[startLockTicker()] defaultIdleReset =", defaultIdleReset)
		if lockTickerCount > 0 {
			tools.Debug("[startLockTicker()] Starting idle timer to lock screen in", settings.LockTimer, "minutes...")
			<-lockTicker.C
			tools.Debug("[startLockTicker()] Timer ended, locking the screen...")
			if settings.LockTimer != 0 {
				tools.LockWindows()
			} else {
				tools.Debug("[startLockTicker()] Lock screen is disabled")
			}
		}
	}
}

func startCountTicker() {
	resetCountTickerTime = time.Duration(defaultIdleReset) * time.Second
	resetCountTicker = time.NewTicker(resetCountTickerTime)
	for {
		<-resetCountTicker.C
		lockTickerCount = defaultIdleReset
		tools.Debug("[startCountTicker()] lockTickerCount reset to", lockTickerCount)
	}
}

func startIdleTicker() {
	idleTicker = time.NewTicker(1 * time.Second)
	for range idleTicker.C {
		if tools.IdleTime() < 1000000000 {
			if lockTickerCount <= defaultIdleReset {
				lockTickerCount--
			}
			if lockTickerCount < defaultIdleReset {
				tools.Debug("[startIdleTicker()] settings.LockTimer =", settings.LockTimer)
				tools.Debug("[startIdleTicker()] Stopping resetCountTicker...")
				resetCountTicker.Reset(resetCountTickerTime)
				if lockTickerCount <= 0 {
					tools.Debug("[startIdleTicker()] Stopping idle timer...")
					lockTicker.Reset(checkLockTimer() * time.Minute)
				}
			}
			tools.Debug(fmt.Sprintf("\n***************Countdown reset at %s ***************\n", time.Now()))
		}
		tools.Debug("[startIdleTicker()] lockTickerCount = ", lockTickerCount)
		tools.Debug("[startIdleTicker()] System idles for", tools.IdleTime())
	}
}

func checkLockTimer() time.Duration {
	if settings.LockTimer > 0 {
		return time.Duration(settings.LockTimer)
	}
	return time.Duration(settings.LockTimer + 1)
}
