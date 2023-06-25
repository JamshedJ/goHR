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

	foldername := configs.Logger.FolderName
	if err = os.Mkdir(foldername, 0777); err != nil && err.Error() != "mkdir logs: file exists" {
		return fmt.Errorf("logger Init: error creating %s directory: %w", foldername, err)
	}

	for l, fileName := range map[logLevel]string{
		Debug:   configs.Logger.FileDebug,
		Info:    configs.Logger.FileInfo,
		Warning: configs.Logger.FileWarning,
		Error:   configs.Logger.FileError,
	} {
		if err = setLevelOutput(l, foldername+"/"+fileName); err != nil {
			return fmt.Errorf("logger Init: error setting logger output for %q level: %w", l, err)
		}
	}
	return
}

func setLevelOutput(l logLevel, fileName string) (err error) {
	_, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    configs.Logger.MaxSizeMB,
		MaxBackups: configs.Logger.MaxBackups,
		MaxAge:     configs.Logger.MaxAgeDays,
		Compress:   configs.Logger.Compress,
		LocalTime:  true,
	}

	level[l] = log.New(lumberjackLogger, "", log.Ldate|log.Lmicroseconds)
	return
}
