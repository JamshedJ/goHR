package log

import (
	"fmt"
	"log"
	"os"

	"github.com/JamshedJ/goHR/internal/configs"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logLevel string

func (l logLevel) Println(v ...interface{}) {
	level[l].Println(v...)
}

func (l logLevel) Printf(format string, v ...interface{}) {
	level[l].Printf(format, v...)
}

func (l logLevel) Fatal(v ...interface{}) {
	level[l].Fatal(v...)
}

const (
	Debug   logLevel = "debug"
	Info    logLevel = "info"
	Warning logLevel = "warning"
	Error   logLevel = "error"
)

var level map[logLevel]*log.Logger

func Init() (err error) {
	level = make(map[logLevel]*log.Logger)

	for l, fileName := range map[logLevel]string{
		Debug:   configs.AppSettings.AppParams.LogDebug,
		Info:    configs.AppSettings.AppParams.LogInfo,
		Warning: configs.AppSettings.AppParams.LogWarning,
		Error:   configs.AppSettings.AppParams.LogError,
	} {
		if err = setLevelOutput(l, fileName); err != nil {
			return fmt.Errorf("logger Init: error setting logger output for %q level: %w", l, err)
		}
	}
	return
}

func setLevelOutput(l logLevel, fileName string) (err error) {
	_, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    configs.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: configs.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     configs.AppSettings.AppParams.LogMaxAge,   // days
		Compress:   configs.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	level[l] = log.New(lumberjackLogger, "", log.Ldate|log.Lmicroseconds)
	return
}
