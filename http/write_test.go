package http

import (
	ghttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gaffatape-io/gopherrs"
)

func TestError(t *testing.T) {
	tests := []struct {
		err    error
		status int
	}{
		{
			gopherrs.InvalidArgument(nil),
			ghttp.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		Error(rec, tc.err)
		t.Log(rec)

		if rec.Code != tc.status {
			t.Fatal()
		}
	}
}
