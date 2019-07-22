package gopherrs

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		code  codes.Code
		ctor  func(opts ...Option) *GRPCError
		check func(err error) bool
	}{
		{
			codes.Canceled,
			Canceled,
			IsCanceled,
		},
		{
			codes.Unknown,
			Unknown,
			IsUnknown,
		},
		{
			codes.InvalidArgument,
			InvalidArgument,
			IsInvalidArgument,
		},
		{
			codes.DeadlineExceeded,
			DeadlineExceeded,
			IsDeadlineExceeded,
		},
		{
			codes.NotFound,
			NotFound,
			IsNotFound,
		},
		{
			codes.AlreadyExists,
			AlreadyExists,
			IsAlreadyExists,
		},
		{
			codes.PermissionDenied,
			PermissionDenied,
			IsPermissionDenied,
		},
		{
			codes.ResourceExhausted,
			ResourceExhausted,
			IsResourceExhausted,
		},
		{
			codes.Aborted,
			Aborted,
			IsAborted,
		},
		{
			codes.OutOfRange,
			OutOfRange,
			IsOutOfRange,
		},
		{
			codes.Unimplemented,
			Unimplemented,
			IsUnimplemeted,
		},
		{
			codes.Internal,
			Internal,
			IsInternal,
		},
		{
			codes.Unavailable,
			Unavailable,
			IsUnavailable,
		},
		{
			codes.DataLoss,
			DataLoss,
			IsDataLoss,
		},
		{
			codes.Unauthenticated,
			Unauthenticated,
			IsUnauthenticated,
		},
	}

	for _, tc := range tests {
		t.Log(tc, "config:", Configuration)
		err := tc.ctor()
		t.Log(err)

		if err.Code != tc.code {
			t.Fatal()
		}

		err2 := tc.ctor(NoStackTrace)
		t.Log(err2)
		if err.Error() == err2.Error() {
			t.Fatal()
		}

		match := isCode(tc.code, err)
		if !match {
			t.Fatal()
		}

		if !tc.check(err) {
			t.Fatal()
		}

		errStatus := status.Error(tc.code, "ignored")
		t.Log(errStatus)
		if !tc.check(errStatus) {
			t.Fatal()
		}
	}
}

func TestIsCodeSuccessShortcut(t *testing.T) {
	ok := isCode(codes.OK, nil)
	if !ok {
		t.Fatal(ok)
	}
}
