package log

import (
	"fmt"
)

const timestampFormat string = "2006/01/02 15:32:05"

type JSONFormatter struct {
}

func (f *JSONFormatter) Format(e *Entry) string {
	str := fmt.Sprintf(`"caller":"%s:%d","status":"%s","time":"%s","message":"%s"`,
		e.caller.File, e.caller.Line, e.status, e.time.Format(timestampFormat), e.message)

	if e.err != nil {
		str += fmt.Sprintf(`,"error":"%v"`, e.err)
	}

	return fmt.Sprintf(`{%s}`, str)
}

type TextFormatter struct {
}

func (f *TextFormatter) Format(e *Entry) string {
	return fmt.Sprintf("%s |%s\n",
		e.time.Format(timestampFormat),
		e.message)
}
