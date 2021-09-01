package gopherrs

import (
	"runtime"
	"strings"
	"testing"
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
			NewE(InvalidArgument, msg),
			"InvalidArgument abcd\nfoo foo.go:1\n",
		},
		{
			NewEf(InvalidArgument, msgf, 1, 2, 3),
			"InvalidArgument 1 2 3\nfoo foo.go:1\n",
		},
		{
			Wrap(NewE(InvalidArgument, msg), Unknown, msg2),
			"Unknown efgh\nfoo foo.go:1\n  InvalidArgument abcd\n  foo foo.go:1\n",
		},
		{
			Wrapf(NewE(InvalidArgument, msg), Unknown, msgf, 4, 5, 6),
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
	e := NewEf(InvalidArgument, "foo %s", "bar")
	t.Log(e)
	
	if !strings.HasPrefix(e.Error(), "InvalidArgument foo bar") {
		t.Fatal()
	}
}

func TestNewE(t *testing.T) {
	e := NewE(InvalidArgument, "bar foo")
	t.Log(e)

	if !strings.HasPrefix(e.Error(), "InvalidArgument bar foo") {
		t.Fatal()
	}
}
