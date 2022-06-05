package tools

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
)

const (
	t1  = 1
	t2  = 5
	t3  = 10
	t4  = 15
	t5  = 20
	t6  = 30
	t7  = 45
	t8  = 60
	t9  = 90
	t10 = 120
)

var (
	configFileName         = "wl_config.ini"
	defaultLockTime uint16 = 10
	defaultConfig          = fmt.Sprintf(`
# Set options of lock screen timer in minutes
# Should have 10 options from t1 to t10 and values from 1 to 1440
[lock_time_options]
t1 = %d
t2 = %d
t3 = %d
t4 = %d
t5 = %d
t6 = %d
t7 = %d
t8 = %d
t9 = %d
t10 = %d

[settings]
# Current screen lock timer
# Should have values from 0 to 1440, 0 disables screen lock
lockTimer = %d
`, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, defaultLockTime)
)

func configNotExists() bool {
	if _, err := os.Stat(configFileName); err == nil {
		return false
	}
	return true
}

func createConfig() error {
	f, err := os.Create(configFileName)
	defer func(f *os.File) {
		err := f.Close()
		IsError(err)
	}(f)

	if IsError(err) {
		return err
	}
	_, err = f.WriteString(defaultConfig)
	if IsError(err) {
		return err
	}
	err = f.Sync()
	if IsError(err) {
		return err
	}
	return nil
}

type LockTimeOptions struct {
	T1  uint16
	T2  uint16
	T3  uint16
	T4  uint16
	T5  uint16
	T6  uint16
	T7  uint16
	T8  uint16
	T9  uint16
	T10 uint16
}

type Settings struct {
	LockTimer uint16
}

func ReadConfig() (LockTimeOptions, Settings, error) {
	var l LockTimeOptions
	var s Settings

	cfg := openConfig()

	s.LockTimer = func() uint16 {
		v, err := cfg.Section("settings").Key("lockTimer").Int()
		if IsError(err) || v > 1440 {
			return defaultLockTime
		}
		if v < 0 {
			return 0
		}
		return uint16(v)
	}()
	cfgT1, err1 := cfg.Section("lock_time_options").Key("t1").Int()
	cfgT2, err2 := cfg.Section("lock_time_options").Key("t2").Int()
	cfgT3, err3 := cfg.Section("lock_time_options").Key("t3").Int()
	cfgT4, err4 := cfg.Section("lock_time_options").Key("t4").Int()
	cfgT5, err5 := cfg.Section("lock_time_options").Key("t5").Int()
	cfgT6, err6 := cfg.Section("lock_time_options").Key("t6").Int()
	cfgT7, err7 := cfg.Section("lock_time_options").Key("t7").Int()
	cfgT8, err8 := cfg.Section("lock_time_options").Key("t8").Int()
	cfgT9, err9 := cfg.Section("lock_time_options").Key("t9").Int()
	cfgT10, err10 := cfg.Section("lock_time_options").Key("t10").Int()

	l.T1 = verifyValues(t1, cfgT1, err1)
	l.T2 = verifyValues(t2, cfgT2, err2)
	l.T3 = verifyValues(t3, cfgT3, err3)
	l.T4 = verifyValues(t4, cfgT4, err4)
	l.T5 = verifyValues(t5, cfgT5, err5)
	l.T6 = verifyValues(t6, cfgT6, err6)
	l.T7 = verifyValues(t7, cfgT7, err7)
	l.T8 = verifyValues(t8, cfgT8, err8)
	l.T9 = verifyValues(t9, cfgT9, err9)
	l.T10 = verifyValues(t10, cfgT10, err10)
	return l, s, nil
}

// verifyValues checks if t1-t10 values from the config file are correct, otherwise sets default
func verifyValues(d, v int, e error) uint16 {
	if v < 1 || v > 1440 || IsError(e) {
		return uint16(d)
	}
	return uint16(v)
}

func SetLockTimer(t uint16, settings *Settings) {
	cfg := openConfig()
	cfg.Section("settings").Key("lockTimer").SetValue(strconv.Itoa(int(t)))
	err := cfg.SaveTo(configFileName)
	IsError(err)
	settings.LockTimer = t
}

func openConfig() *ini.File {
	if configNotExists() {
		err := createConfig()
		if IsError(err) {
			return nil
		}
	}
	cfg, err := ini.Load(configFileName)
	if IsError(err) {
		return nil
	}
	return cfg
}
