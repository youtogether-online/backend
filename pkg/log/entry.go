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
	caller  *runtime.Frame
	message string
	status  string
	err     error
}

func newEntry(l *Logger, status string) *Entry {
	return &Entry{l: l, status: status}
}

var m sync.Mutex

func (e *Entry) send(args ...any) {
	e.time = time.Now()
	e.message = fmt.Sprint(args...)
	e.printStr()
}

func (e *Entry) sendf(format string, args ...any) {
	e.time = time.Now()
	e.message = fmt.Sprintf(format, args...)
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
		e.status = debug
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (e *Entry) Debugf(format string, args ...any) {
	if e.l.level >= DebugLevel {
		e.status = debug
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (e *Entry) Info(args ...any) {
	if e.l.level >= InfoLevel {
		e.status = info
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (e *Entry) Infof(format string, args ...any) {
	if e.l.level >= InfoLevel {
		e.status = info
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (e *Entry) Warn(args ...any) {
	if e.l.level >= WarnLevel {
		e.status = warning
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (e *Entry) Warnf(format string, args ...any) {
	if e.l.level >= WarnLevel {
		e.status = warning
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (e *Entry) Err(args ...any) {
	if e.l.level >= ErrLevel {
		e.status = err
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.send(args...)
	}
}

func (e *Entry) Errf(format string, args ...any) {
	if e.l.level >= ErrLevel {
		e.status = err
		if e.l.reportCaller {
			e.caller = getReportCaller()
		}
		go e.sendf(format, args...)
	}
}

func (e *Entry) Fatal(args ...any) {
	e.status = fatal
	if e.l.reportCaller {
		e.caller = getReportCaller()
	}
	e.send(args...)
	os.Exit(1)
}

func (e *Entry) Fatalf(format string, args ...any) {
	e.status = fatal
	if e.l.reportCaller {
		e.caller = getReportCaller()
	}
	e.sendf(format, args...)
	os.Exit(1)
}
