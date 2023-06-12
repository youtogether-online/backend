package log

var logger = NewLogger(InfoLevel, &JSONFormatter{}, true)

func SetLevel(level Level) {
	logger.level = level
}

func WithErr(err error) *Entry {
	return logger.WithErr(err)
}

func Debug(args ...any) {
	logger.Debug(args...)
}

func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

func Info(args ...any) {
	logger.Info(args...)
}

// LastInfo starts without goroutine
func LastInfo(args ...any) {
	if logger.level >= InfoLevel {
		e := newEntry(logger, info)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		e.send(args...)
	}
}

func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

func Warn(args ...any) {
	logger.Warn(args...)
}

func Warnf(format string, args ...any) {
	logger.Warnf(format, args...)
}

func Err(args ...any) {
	logger.Err(args...)
}

func Errf(format string, args ...any) {
	logger.Errf(format, args...)
}

func Fatal(args ...any) {
	logger.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}
