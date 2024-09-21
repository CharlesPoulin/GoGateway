package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	// Add other methods as needed
}

type logrusLogger struct {
	*logrus.Logger
}

func NewLogger(level string) Logger {
	logger := logrus.New()
	logger.Out = os.Stdout

	// Set log level
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(lvl)
	}

	// Set formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z07:00",
	})

	return &logrusLogger{logger}
}
