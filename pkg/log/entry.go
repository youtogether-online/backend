package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

type Entry struct {
	l       *Logger
	time    time.Time
	caller  string
	message string
	status  Level
	err     string
}

var m sync.Mutex

func fillEntry(e *Entry, msg string) {
	e.message = msg
	if e.l.reportCaller {
		_, file, line, ok := runtime.Caller(3)
		if !ok {
			panic("can't get caller")
		}
		e.caller = fmt.Sprintf("%s#%d", file, line)
	}
	e.printStr()
}

func (e *Entry) printStr() {
	str := e.l.formatter.Format(e)

	m.Lock()
	defer m.Unlock()
	fmt.Println(str)
}

func (e *Entry) Debug(args ...any) {
	if e.l.level >= DebugLevel {
		go send(e, args...)
	}
}

func (e *Entry) Debugf(format string, args ...any) {
	if e.l.level >= DebugLevel {
		go sendf(e, format, args...)
	}
}

func (e *Entry) Info(args ...any) {
	if e.l.level >= InfoLevel {
		go send(e, args...)
	}
}

func (e *Entry) Infof(format string, args ...any) {
	if e.l.level >= InfoLevel {
		go sendf(e, format, args...)
	}
}

func (e *Entry) Warn(args ...any) {
	if e.l.level >= WarnLevel {
		go send(e, args...)
	}
}

func (e *Entry) Warnf(format string, args ...any) {
	if e.l.level >= WarnLevel {
		go sendf(e, format, args...)
	}
}

func (e *Entry) Err(args ...any) {
	if e.l.level >= ErrLevel {
		go send(e, args...)
	}
}

func (e *Entry) Errf(format string, args ...any) {
	if e.l.level >= ErrLevel {
		go sendf(e, format, args...)
	}
}

func (e *Entry) Fatal(args ...any) {
	if e.l.level == FatalLevel {
		send(e, args...)
		os.Exit(1)
	}
}

func (e *Entry) Fatalf(format string, args ...any) {
	if e.l.level == FatalLevel {
		sendf(e, format, args...)
		os.Exit(1)
	}
}
