package http

import (
	ghttp "net/http"

	"github.com/gaffatape-io/gopherrs"
	"google.golang.org/grpc/codes"
)

var (
	// mapping based on: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
	code2status = map[codes.Code]int{
		codes.OK: ghttp.StatusOK,
		// 499 is non-standard nginx error code
		codes.Canceled:           499,
		codes.Unknown:            ghttp.StatusInternalServerError,
		codes.InvalidArgument:    ghttp.StatusBadRequest,
		codes.DeadlineExceeded:   ghttp.StatusGatewayTimeout,
		codes.NotFound:           ghttp.StatusNotFound,
		codes.AlreadyExists:      ghttp.StatusConflict,
		codes.PermissionDenied:   ghttp.StatusForbidden,
		codes.Unauthenticated:    ghttp.StatusUnauthorized,
		codes.ResourceExhausted:  ghttp.StatusTooManyRequests,
		codes.FailedPrecondition: ghttp.StatusBadRequest,
		codes.Aborted:            ghttp.StatusConflict,
		codes.OutOfRange:         ghttp.StatusBadRequest,
		codes.Unimplemented:      ghttp.StatusNotImplemented,
		codes.Internal:           ghttp.StatusInternalServerError,
		codes.Unavailable:        ghttp.StatusServiceUnavailable,
		codes.DataLoss:           ghttp.StatusInternalServerError,
	}
)

// Error replaces the http.Error() function in the http package.
func Error(w ghttp.ResponseWriter, err error) {
	status := code2status[gopherrs.Code(err)]
	ghttp.Error(w, err.Error(), status)
}
