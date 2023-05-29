package log

import (
	"fmt"
)

var timestampFormat = "2006/01/02 15:32:05"

type JSONFormatter struct {
}

func (f *JSONFormatter) Format(e *Entry) string {
	str := fmt.Sprintf(`"caller":"%s","status":"%d","time":"%s","message":"%s"`,
		e.caller, e.status, e.time.Format(timestampFormat), e.message)

	if e.err != "" {
		str += fmt.Sprintf(`,"error":"%s"`, e.err)
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
