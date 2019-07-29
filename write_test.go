package gopherrs

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		err    error
		status int
	}{
		{
			InvalidArgument(nil),
			http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		Write(rec, tc.err)
		t.Log(rec)

		if rec.Code != tc.status {
			t.Fatal()
		}
	}
}
