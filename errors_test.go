package gopherrs

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/gaffatape-io/gopherrs/codes"
)

func TestEErrorCallstack(t *testing.T) {
	const msg = "abcd"
	const msg2 = "efgh"
	const msgf = "%d %d %d"
	tests := []struct {
		err *E
		txt string
	} {
		{
			NewE(codes.InvalidArgument, msg),
			"InvalidArgument abcd\nfoo foo.go:1\n",
		},
		{
			NewEf(codes.InvalidArgument, msgf, 1, 2, 3),
			"InvalidArgument 1 2 3\nfoo foo.go:1\n",
		},
		{
			Wrap(NewE(codes.InvalidArgument, msg), codes.Unknown, msg2),
			"Unknown efgh\nfoo foo.go:1\n  InvalidArgument abcd\n  foo foo.go:1\n",
		},
		{
			Wrapf(NewE(codes.InvalidArgument, msg), codes.Unknown, msgf, 4, 5, 6),
			"Unknown 4 5 6\nfoo foo.go:1\n  InvalidArgument abcd\n  foo foo.go:1\n",
		},
		
	}

	callersFrames = func([]uintptr) framesIter {
		return &fakeFrames{testFrames, 0}
	}
	
	for _, tc := range tests {
		txt := tc.err.Error()
		t.Logf("g:%q", txt)
		t.Logf("w:%q", tc.txt)

		if txt != tc.txt {
			t.Fatal()
		}
	}
}

type fakeFrames struct {
	frames []runtime.Frame
	at int
}

func (f *fakeFrames) Next() (runtime.Frame, bool) {
	if f.at == len(f.frames) {
		return runtime.Frame{}, false
	}

	cur := f.frames[f.at]
	f.at++
	return cur, f.at == len(f.frames)
}

var (
	testFrames = []runtime.Frame{
		runtime.Frame{Function: "foo", File: "foo.go", Line: 1},
		runtime.Frame{Function: "bar", File: "bar.go", Line: 2},
		runtime.Frame{Function: "baz", File: "baz.go", Line: 3},
	}
)


func TestNewEf(t *testing.T) {
	e := NewEf(codes.InvalidArgument, "foo %s", "bar")
	t.Log(e)
	
	if !strings.HasPrefix(e.Error(), "InvalidArgument foo bar") {
		t.Fatal()
	}
}

func TestNewE(t *testing.T) {
	e := NewE(codes.InvalidArgument, "bar foo")
	t.Log(e)

	if !strings.HasPrefix(e.Error(), "InvalidArgument bar foo") {
		t.Fatal()
	}
}

func TestErrorsIs(t *testing.T) {
	tests := []struct {
		err *E
		c codes.Code
		is bool
	} {
		{
			NewE(codes.InvalidArgument, ""),
			codes.InvalidArgument,
			true,
		},
		{	
			Wrap(NewE(codes.InvalidArgument, ""), codes.Unknown, ""),
			codes.Unknown,
			true,
		},
	}

	for _, tc := range tests {
		g := errors.Is(tc.err, tc.c)
		t.Log(tc)
		if g != tc.is {
			t.Fatal()
		}
	}
}


func TestGenerateCodeCtors(t *testing.T) {
	for _, c := range []codes.Code{
		codes.Canceled,
		codes.Unknown,
		codes.InvalidArgument,
		codes.DeadlineExceeded,
		codes.PermissionDenied,
		codes.ResourceExhausted,
		codes.FailedPrecondition,
		codes.Aborted,
		codes.OutOfRange,
		codes.Unimplemented,
		codes.Internal,
		codes.Unavailable,
		codes.DataLoss,
		codes.Unauthenticated,
	} {
		fmt.Printf("func %s(msg string) *E {\n", c)
		fmt.Printf("\treturn NewE(codes.%s, msg)\n", c)
		fmt.Println("}")

		fmt.Printf("func %sf(msg string, args ...interface{}) *E {\n", c)
		fmt.Printf("\treturn NewEf(codes.%s, msg, args...)\n", c)
		fmt.Println("}")

		fmt.Printf("func Wrap%s(cause error, msg string) *E {\n", c)
		fmt.Printf("\treturn Wrap(cause, codes.%s, msg)\n", c)
		fmt.Println("}")

		fmt.Printf("func Wrap%sf(cause error, msg string, args ...interface{}) *E {\n", c)
		fmt.Printf("\treturn Wrapf(cause, codes.%s, msg, args...)\n", c)
		fmt.Println("}")
	}

	
}

