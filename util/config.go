package util

import (
	"strings"
)

// Load read data from section in ini file
func LoadSection(section string) {
	AppEnv.Config.Section(section)
}

func LoadList(section, key string) (slice []string) {
	cfg := AppEnv.Config
	slice = strings.Split(cfg.Section(section).Key(key).String(), ",")
	return
}

// Getkey ...
func Getkey(section, key string) (skey string) {
	cfg := AppEnv.Config
	AppEnv.Log.Debug(cfg.Section("server").Key("httpPort").String())
	skey = cfg.Section(section).Key(key).String()
	return
}
