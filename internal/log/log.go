package log

import "github.com/sirupsen/logrus"

var loggerInstance *logrus.Logger

func init() {
	loggerInstance = NewLogger()
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	return logger
}

func Logger() *logrus.Logger {
	if loggerInstance == nil {
		loggerInstance = NewLogger()
	}

	return loggerInstance
}
