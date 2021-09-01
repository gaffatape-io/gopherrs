package gopherrs

import (
	"flag"
	"fmt"
	"strings"
	"runtime"
)

var maxCallstackDepth = flag.Int("max_callstack_depth", 32, "Max stackframes to capture when an error is created")

// E is the implementation of the error interface.
type E struct {
	Code Code
	Message string
	FormatArgs []interface{}
	Callstack []uintptr
}

func callstack() []uintptr {
	s := make([]uintptr, *maxCallstackDepth)
	n := runtime.Callers(3, s)
	return s[:n]
}

// NewE returns a newly created error.
func NewE(code Code, msg string) *E {
	return &E{code, msg, nil, callstack()}
}

// NewEf returns a new created error with formatting.
func NewEf(code Code, msg string, args ...interface{}) *E {
	return &E{code, msg, args, callstack()}
}

func writeCallstack(b *strings.Builder, callstack []uintptr) {
	frames := runtime.CallersFrames(callstack)
	for {
		frame, ok := frames.Next()
		b.WriteString(frame.Function)
		b.WriteByte(' ')
		b.WriteString(frame.File)
		b.WriteByte(':')
		fmt.Fprint(b, frame.Line)
		b.WriteByte('\n')
		
		if !ok {
			return
		}
	}
}

// Error implements the error interface.
func (e *E) Error() string {
	b := &strings.Builder{}
	b.WriteString(e.Code.String())
	b.WriteByte(' ')
	if e.FormatArgs == nil {
		b.WriteString(e.Message)
	} else {
		b.WriteString(fmt.Sprintf(e.Message, e.FormatArgs...))
	}

	b.WriteByte('\n')
	writeCallstack(b, e.Callstack)
	return b.String()
}
