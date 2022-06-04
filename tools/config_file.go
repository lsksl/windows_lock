package tools

import (
	"gopkg.in/ini.v1"
	"os"
)

var (
	defaultConfig = `
# Set options of lock screen timer in minutes
# Should have 10 options from t1 to t10
[lock_time_options]
t1 = 5
t2 = 8
t3 = 10
t4 = 12
t5 = 15
t6 = 30
t7 = 45
t8 = 60
t9 = 90
t10 = 120

# Current screen lock timer
[settings]
lockTimer = 5
`
	configFileName = "wl_config.ini"
)

func configNotExists() bool {
	if _, err := os.Stat(configFileName); err == nil {
		return false
	}
	return true
}

func createConfig() error {
	f, err := os.Create(configFileName)
	defer f.Close()
	if IsError(err) {
		return err
	}
	_, err = f.WriteString(defaultConfig)
	if IsError(err) {
		return err
	}
	f.Sync()
	return nil
}

type LockTimeOptions struct {
	T1  uint
	T2  uint
	T3  uint
	T4  uint
	T5  uint
	T6  uint
	T7  uint
	T8  uint
	T9  uint
	T10 uint
}

type Settings struct {
	LockTimer uint
}

func ReadConfig() (LockTimeOptions, Settings, error) {
	var l LockTimeOptions
	var s Settings

	if configNotExists() {
		err := createConfig()
		if IsError(err) {
			return l, s, err
		}
	}
	cfg, err := ini.Load(configFileName)
	if IsError(err) {
		return l, s, err
	}
	s.LockTimer, _ = cfg.Section("settings").Key("lockTimer").Uint()
	l.T1, _ = cfg.Section("lock_time_options").Key("t1").Uint()
	l.T2, _ = cfg.Section("lock_time_options").Key("t2").Uint()
	l.T3, _ = cfg.Section("lock_time_options").Key("t3").Uint()
	l.T4, _ = cfg.Section("lock_time_options").Key("t4").Uint()
	l.T5, _ = cfg.Section("lock_time_options").Key("t5").Uint()
	l.T6, _ = cfg.Section("lock_time_options").Key("t6").Uint()
	l.T7, _ = cfg.Section("lock_time_options").Key("t7").Uint()
	l.T8, _ = cfg.Section("lock_time_options").Key("t8").Uint()
	l.T9, _ = cfg.Section("lock_time_options").Key("t9").Uint()
	l.T10, _ = cfg.Section("lock_time_options").Key("t10").Uint()
	return l, s, nil
}
