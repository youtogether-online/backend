package log

import (
	"os"
)

var logger = NewLogger(InfoLevel, &JSONFormatter{}, true)

func SetLevel(level Level) {
	logger.level = level
}

func WithErr(err error) *Entry {
	return &Entry{
		l:   logger,
		err: err,
	}
}

func Debug(args ...any) {
	if logger.level >= DebugLevel {
		e := newEntry(logger, debug)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func Debugf(format string, args ...any) {
	if logger.level >= DebugLevel {
		e := newEntry(logger, debug)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func Info(args ...any) {
	if logger.level >= InfoLevel {
		e := newEntry(logger, info)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func Infof(format string, args ...any) {
	if logger.level >= InfoLevel {
		e := newEntry(logger, info)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func Warn(args ...any) {
	if logger.level >= WarnLevel {
		e := newEntry(logger, warning)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func Warnf(format string, args ...any) {
	if logger.level >= WarnLevel {
		e := newEntry(logger, warning)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func Err(args ...any) {
	if logger.level >= ErrLevel {
		e := newEntry(logger, err)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func Errf(format string, args ...any) {
	if logger.level >= ErrLevel {
		e := newEntry(logger, err)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func Fatal(args ...any) {
	if logger.level == FatalLevel {
		e := newEntry(logger, fatal)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		e.send(args...)
		os.Exit(1)
	}
}

func Fatalf(format string, args ...any) {
	if logger.level == FatalLevel {
		e := newEntry(logger, fatal)
		if logger.reportCaller {
			e.caller = getReportCaller()
		}
		e.sendf(format, args...)
		os.Exit(1)
	}
}
