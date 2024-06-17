package log

import "github.com/sirupsen/logrus"

var loggerInstance *logrus.Logger

func init() {
	loggerInstance = Init()
}

func Init() *logrus.Logger {
	logger := logrus.New()

	return logger
}

func Logger() *logrus.Logger {
	if loggerInstance == nil {
		Init()
	}

	return loggerInstance
}
