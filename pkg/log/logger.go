package log

import (
	"os"
	"runtime"
)

type Level uint16

const (
	FatalLevel Level = iota
	ErrLevel
	WarnLevel
	InfoLevel
	DebugLevel
	fatal   string = "fatal"
	err     string = "error"
	warning string = "warning"
	info    string = "info"
	debug   string = "debug"
)

type Formatter interface {
	Format(*Entry) string
}

type Logger struct {
	level        Level
	formatter    Formatter
	reportCaller bool
}

func NewLogger(level Level, formatter Formatter, reportCaller bool) *Logger {
	return &Logger{level: level, formatter: formatter, reportCaller: reportCaller}
}

const frameIndex int = 3

func getReportCaller() *runtime.Frame {

	programCounters := make([]uintptr, frameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, i := true, 0; more && i <= frameIndex; i++ {
			if i == frameIndex {
				frame, more = frames.Next()
			} else {
				_, more = frames.Next()
			}
		}
	}
	return &frame
}

func (l *Logger) WithErr(err error) *Entry {
	if err == nil {
		return &Entry{
			l: l,
		}
	}
	return &Entry{
		l:   l,
		err: err,
	}
}

func (l *Logger) Debug(args ...any) {
	if l.level >= DebugLevel {
		e := newEntry(l, debug)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.level >= DebugLevel {
		e := newEntry(l, debug)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (l *Logger) Info(args ...any) {
	if l.level >= InfoLevel {
		e := newEntry(l, info)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	if l.level >= InfoLevel {
		e := newEntry(l, info)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (l *Logger) Warn(args ...any) {
	if l.level >= WarnLevel {
		e := newEntry(l, warning)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.level >= WarnLevel {
		e := newEntry(l, warning)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (l *Logger) Err(args ...any) {
	if l.level >= ErrLevel {
		e := newEntry(l, err)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (l *Logger) Errf(format string, args ...any) {
	if l.level >= ErrLevel {
		e := newEntry(l, err)
		if l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (l *Logger) Fatal(args ...any) {
	e := newEntry(l, fatal)
	if l.reportCaller {
		e.caller = getReportCaller()
	}
	e.send(args...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...any) {
	e := newEntry(l, fatal)
	if l.reportCaller {
		e.caller = getReportCaller()
	}
	e.sendf(format, args...)
	os.Exit(1)
}
