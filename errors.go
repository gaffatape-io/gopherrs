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
	Message string
	Callstack []uintptr
}

func callstack() []uintptr {
	s := make([]uintptr, *maxCallstackDepth)
	n := runtime.Callers(3, s)
	return s[:n]
}

func NewE(msg string) *E {
	return &E{msg, callstack()}
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
	b.WriteString(e.Message)
	b.WriteByte('\n')
	writeCallstack(b, e.Callstack)
	return b.String()
}
