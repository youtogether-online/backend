package log

import "os"

var logger = &Logger{
	reportCaller: true,
	level:        InfoLevel,
	formatter:    &JSONFormatter{},
}

func SetLevel(level Level) {
	logger.level = level
}

func WithErr(err error) *Entry {
	return &Entry{
		l:   logger,
		err: err.Error(),
	}
}

func Debug(args ...any) {
	if logger.level >= DebugLevel {
		go send(&Entry{l: logger, status: DebugLevel}, args...)
	}
}

func Debugf(format string, args ...any) {
	if logger.level >= DebugLevel {
		go sendf(&Entry{l: logger, status: DebugLevel}, format, args...)
	}
}

func Info(args ...any) {
	if logger.level >= InfoLevel {
		go send(&Entry{l: logger, status: InfoLevel}, args...)
	}
}

func Infof(format string, args ...any) {
	if logger.level >= InfoLevel {
		go sendf(&Entry{l: logger, status: InfoLevel}, format, args...)
	}
}

func Warn(args ...any) {
	if logger.level >= WarnLevel {
		go send(&Entry{l: logger, status: WarnLevel}, args...)
	}
}

func Warnf(format string, args ...any) {
	if logger.level >= WarnLevel {
		go sendf(&Entry{l: logger, status: WarnLevel}, format, args...)
	}
}

func Err(args ...any) {
	if logger.level >= ErrLevel {
		go send(&Entry{l: logger, status: ErrLevel}, args...)
	}
}

func Errf(format string, args ...any) {
	if logger.level >= ErrLevel {
		go sendf(&Entry{l: logger, status: ErrLevel}, format, args...)
	}
}

func Fatal(args ...any) {
	if logger.level == FatalLevel {
		send(&Entry{l: logger, status: FatalLevel}, args...)
		os.Exit(1)
	}
}

func Fatalf(format string, args ...any) {
	if logger.level == FatalLevel {
		go sendf(&Entry{l: logger, status: FatalLevel}, format, args...)
		os.Exit(1)
	}
}
