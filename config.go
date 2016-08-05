package logger

import (
	"os"
	"github.com/kyugao/gy.util.json"
)

var Config struct {
	LogPath       string
	LogFile       string
	EnableConsole bool
	Level         string
}

func init() {
	err := ujson.FromFile("conf/logger.json", &Config)
	if err != nil {
		Error("load logger config error %s.\n", ujson.ToJsonString(err))
		panic(err)
	}

	_, err = os.Stat(Config.LogPath)
	if !os.IsExist(err) {
		Debug("Log path missing, create.")
		err = os.MkdirAll(Config.LogPath, os.ModePerm)
		if err != nil {
			Fatal("Create log path error", err)
			panic(err)
		}
	} else {
		Debug("Log path exists, skip.")
	}

	var level LEVEL
	switch Config.Level {
	case "ALL":
		level = ALL
	case "DEBUG":
		level = DEBUG
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	case "FATAL":
		level = FATAL
	case "OFF":
		level = OFF
	}
	SetLevel(level)
	SetConsole(Config.EnableConsole)
	SetRollingDaily(Config.LogPath, Config.LogFile)
}