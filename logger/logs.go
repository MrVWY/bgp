package logger

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var Logger *logrus.Entry

const (
	Path = "/root/"
	Level = "debug"
)

func InitLogger(path, level string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0775)
			if err != nil {
				panic("Failed to create log dir: " + err.Error())
			}
		} else {
			panic("Check log dir failed: " + err.Error())
		}
	}

	logf, err := rotatelogs.New(
		path + "/%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(15 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)
	if err != nil {
		panic("Failed to create rotatelogs: " + err.Error())
	}

	l := logrus.New()
	l.Out = logf

	switch strings.ToLower(level) {
	default:
		fallthrough
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "warn":
		fallthrough
	case "warning":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	}

	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})

	hostname, err := os.Hostname()
	if err != nil {
		panic("Get hostname failed: " + err.Error())
	}

	Logger = l.WithField("hostname", hostname)

	Logger.Info("Logger initializing")
	Logger.Info("Log level: ", strings.ToLower(level))
	Logger.Info("Logger initialize succeed")
}