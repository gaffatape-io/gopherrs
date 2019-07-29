package gopherrs

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	// mapping based on: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
	// We don't use the status definitions from the http package to avoid the dependency.
	code2status = map[codes.Code]int{
		codes.OK:                 100,
		codes.Canceled:           499, // 499 is non-standard nginx error code
		codes.Unknown:            500,
		codes.InvalidArgument:    400,
		codes.DeadlineExceeded:   408,
		codes.NotFound:           404,
		codes.AlreadyExists:      409,
		codes.PermissionDenied:   403,
		codes.Unauthenticated:    401,
		codes.ResourceExhausted:  429,
		codes.FailedPrecondition: 400,
		codes.Aborted:            409,
		codes.OutOfRange:         400,
		codes.Unimplemented:      501,
		codes.Internal:           500,
		codes.Unavailable:        503,
		codes.DataLoss:           500,
	}
)

// Error replaces the http.Error() function in the http package.
func Write(w http.ResponseWriter, err error) {
	status := code2status[Code(err)]
	http.Error(w, err.Error(), status)
}
