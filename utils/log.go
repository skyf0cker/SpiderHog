package utils

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"time"
)

var once sync.Once
var Log *logrus.Logger

func GetFileLogger() *logrus.Logger {
	once.Do(func() {
		logFilePath := "./log"
		logFileName := "Spider.log"
		filename := path.Join(logFilePath, logFileName)

		src, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic("log file opened failed!")
		}

		Log = logrus.New()

		Log.Out = src
		Log.SetLevel(logrus.DebugLevel)

		// 设置 rotatelogs
		logWriter, err := rotatelogs.New(
			// 分割后的文件名称
			filename + ".%Y%m%d.log",

			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(logFileName),

			// 设置最大保存时间(7天)
			rotatelogs.WithMaxAge(7*24*time.Hour),

			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 新增 Hook
		Log.AddHook(lfHook)
	})

	return Log
}

func GetStdoutLogger() *logrus.Logger {
	once.Do(func() {
		Log = logrus.New()

		Log.Out = os.Stdout
		Log.SetLevel(logrus.DebugLevel)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  os.Stdout,
			logrus.FatalLevel: os.Stdout,
			logrus.DebugLevel: os.Stdout,
			logrus.WarnLevel:  os.Stdout,
			logrus.ErrorLevel: os.Stdout,
			logrus.PanicLevel: os.Stdout,
		}
		//
		lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
		})

		//新增 Hook
		Log.AddHook(lfHook)
	})

	return Log
}

