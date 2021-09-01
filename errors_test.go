package gopherrs

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestEError(t *testing.T) {
	b:= &strings.Builder{}	
	s := make([]uintptr, 32)

	msg := "abcd"
	
	// Keep together since this test uses the line number from
	// the first call to figure out the line number of the second call.
	n := runtime.Callers(1, s)
	e := NewE(msg)
	
	writeCallstack(b, s[:n])

	diff := &strings.Builder{}
	got := e.Error()
	want := msg + "\n" + b.String()

	count := len(got)
	if len(want) > len(got) {
		count = len(want)
	}
	t.Log(len(got), len(want))

	// Worst diff ever :) produces a string that can be compared against
	// the expected result, the string have no actual meaning and can be
	// considered a 'fingerprint' of the diff.
	for i := 0; i < count; i++ {
		var g, w byte
		if i < len(got) {
			g = got[i]
		}

		if i < len(want) {
			w = want[i]
		}

		markDiff := func(txt string, i int) string {
			return fmt.Sprintf("%s[%c]%s",txt[i-3:i], txt[i], txt[i+1:i+3])
		}
		
		if g != w {
			t.Logf(fmt.Sprintf("p: %d g:%q w:%q", i, markDiff(got, i), markDiff(want, i)))
		
			diff.WriteByte(g)
			diff.WriteByte(w)
		}
	}

	t.Log("diff:", diff)
	// This check will need to be adjusted as additional imports are added to this file.
	if diff.String() != "98" {
		t.Fatal()
	}
}

