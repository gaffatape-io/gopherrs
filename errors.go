package gopherrs

import (
	"flag"
	"fmt"
	"strings"
	"runtime"
)

type framesIter interface {
	Next() (runtime.Frame, bool)
}

type callersFramesFunc func([]uintptr) framesIter

var (
	maxCallstackDepth = flag.Int("max_callstack_depth", 32, "Max stackframes to capture when an error is created")

	callersFrames callersFramesFunc = func(frames []uintptr) framesIter {
		return runtime.CallersFrames(frames)
	}
)

// E is the implementation of the error interface.
type E struct {
	Code Code
	Message string
	FormatArgs []interface{}
	Callstack []uintptr
	Cause error
}

func callstack() []uintptr {
	s := make([]uintptr, *maxCallstackDepth)
	n := runtime.Callers(3, s)
	return s[:n]
}

// NewE returns a newly created error.
func NewE(code Code, msg string) *E {
	return &E{code, msg, nil, callstack(), nil}
}

// NewEf returns a new created error with formatting.
func NewEf(code Code, msg string, args ...interface{}) *E {
	return &E{code, msg, args, callstack(), nil}
}

// Wrap returns an error that wraps an existing error.
func Wrap(cause error, code Code, msg string) *E {
	return &E{code, msg, nil, callstack(), cause}
}

// Wrapf returns an error with formatting that wraps an existing error.
func Wrapf(cause error, code Code, msg string, args ...interface{}) *E {
	return &E{code, msg, args, callstack(), cause}	
}

func writeCallstack(b *strings.Builder, callstack []uintptr, indent string) {
	frames := callersFrames(callstack)
	for {
		frame, ok := frames.Next()
		b.WriteString(indent)
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
	e.writeError(b, "")
	return b.String()
}

func (e *E) writeError(b *strings.Builder, indent string) {
	b.WriteString(indent)
	b.WriteString(e.Code.String())
	
	b.WriteByte(' ')
	
	if e.FormatArgs == nil {
		b.WriteString(e.Message)
	} else {
		b.WriteString(fmt.Sprintf(e.Message, e.FormatArgs...))
	}

	b.WriteByte('\n')
	
	writeCallstack(b, e.Callstack, indent)

	switch cause := e.Cause.(type) {
	case *E:
		// TODO(dape): replace the indent growth with a fixed set of string to avoid alloc.
		cause.writeError(b, indent + "  ")

	case error:
		b.WriteString(cause.Error())
	}
}

// Unwrap to work well with the errors package.
func (e *E) Unwrap() error {
	return e.Cause
}
