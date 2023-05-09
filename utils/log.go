package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"strings"
)

type LogConf struct {
	Level   string `yaml:"level"`
	Path    string `yaml:"path"`
	Size    int    `yaml:"size"`
	Backups int    `yaml:"backups"`
}

var levelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel, //
	"error": logrus.ErrorLevel,
	"panic": logrus.PanicLevel,
	"trace": logrus.TraceLevel,
	"fatal": logrus.FatalLevel,
}

var Logger *logrus.Entry

func getLogLevel(conf string) logrus.Level {
	if level, ok := levelMap[conf]; ok {
		return level
	}
	// Default level is Info
	return logrus.InfoLevel
}

func setFormatter() logrus.Formatter {
	formatter := &logrus.JSONFormatter{
		//TimestampFormat: "2006-01-02 15:04:05",
		DisableTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			name := frame.Function
			if names := strings.SplitN(frame.Function, "/", 2); len(names) > 1 {
				name = names[1]
			}
			return fmt.Sprintf("%s:%d", name, frame.Line), ""
		},
	}
	return formatter
}

func InitLogger(cfg *LogConf) {
	rotateLogger := &lumberjack.Logger{
		Filename:   cfg.Path,
		MaxSize:    cfg.Size,
		MaxBackups: cfg.Backups,
		LocalTime:  true,
	}
	ioWriter := io.MultiWriter(rotateLogger, os.Stdout)
	logrus.SetOutput(ioWriter)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(setFormatter())
	logrus.SetLevel(getLogLevel(cfg.Level))
	Logger = logrus.WithFields(logrus.Fields{})
	return
}
