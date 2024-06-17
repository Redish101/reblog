package log

func Info(args ...any) {
	Logger().Info(args...)
}

func Debug(args ...any) {
	Logger().Debug(args...)
}

func Warn(args ...any) {
	Logger().Warn(args...)
}

func Error(args ...any) {
	Logger().Error(args...)
}

func Fatal(args ...any) {
	Logger().Fatal(args...)
}

func Panic(args ...any) {
	Logger().Panic(args...)
}

func Infof(format string, args ...any) {
	Logger().Infof(format, args...)
}

func Debugf(format string, args ...any) {
	Logger().Debugf(format, args...)
}

func Warnf(format string, args ...any) {
	Logger().Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	Logger().Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	Logger().Fatalf(format, args...)
}

func Panicf(format string, args ...any) {
	Logger().Panicf(format, args...)
}
