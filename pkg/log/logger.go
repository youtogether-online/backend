package log

import (
	"fmt"
	"os"
	"time"
)

type Level uint16

const (
	FatalLevel Level = iota
	ErrLevel
	WarnLevel
	InfoLevel
	DebugLevel
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

func send(e *Entry, args ...any) {
	e.time = time.Now()
	fillEntry(e, fmt.Sprint(args...))
}

func sendf(e *Entry, format string, args ...any) {
	e.time = time.Now()
	fillEntry(e, fmt.Sprintf(format, args...))
}

func (l *Logger) WithErr(err error) *Entry {
	return &Entry{
		l:   l,
		err: err.Error(),
	}
}

func (l *Logger) Debug(args ...any) {
	if l.level >= DebugLevel {
		go send(&Entry{l: l, status: DebugLevel}, args...)
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.level >= DebugLevel {
		go sendf(&Entry{l: l, status: DebugLevel}, format, args...)
	}
}

func (l *Logger) Info(args ...any) {
	if l.level >= InfoLevel {
		go send(&Entry{l: l, status: InfoLevel}, args...)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	if l.level >= InfoLevel {
		go sendf(&Entry{l: l, status: InfoLevel}, format, args...)
	}
}

func (l *Logger) Warn(args ...any) {
	if l.level >= WarnLevel {
		go send(&Entry{l: l, status: WarnLevel}, args...)
	}
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.level >= WarnLevel {
		go sendf(&Entry{l: l, status: WarnLevel}, format, args...)
	}
}

func (l *Logger) Err(args ...any) {
	if l.level >= ErrLevel {
		go send(&Entry{l: l, status: ErrLevel}, args...)
	}
}

func (l *Logger) Errf(format string, args ...any) {
	if l.level >= ErrLevel {
		go sendf(&Entry{l: l, status: ErrLevel}, format, args...)
	}
}

func (l *Logger) Fatal(args ...any) {
	if l.level == FatalLevel {
		send(&Entry{l: l, status: FatalLevel}, args...)
		os.Exit(1)
	}
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.level == FatalLevel {
		go sendf(&Entry{l: l, status: FatalLevel}, format, args...)
		os.Exit(1)
	}
}
