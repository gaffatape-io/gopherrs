package gopherrs

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		code  codes.Code
		ctor  func(cause error) *GRPCError
		ctorf func(cause error, f string, data ...interface{}) *GRPCError
		check func(err error) bool
	}{
		{
			codes.Canceled,
			Canceled,
			Canceledf,
			IsCanceled,
		},
		{
			codes.Unknown,
			Unknown,
			Unknownf,
			IsUnknown,
		},
		{
			codes.InvalidArgument,
			InvalidArgument,
			InvalidArgumentf,
			IsInvalidArgument,
		},
		{
			codes.DeadlineExceeded,
			DeadlineExceeded,
			DeadlineExceededf,
			IsDeadlineExceeded,
		},
		{
			codes.NotFound,
			NotFound,
			NotFoundf,
			IsNotFound,
		},
		{
			codes.AlreadyExists,
			AlreadyExists,
			AlreadyExistsf,
			IsAlreadyExists,
		},
		{
			codes.PermissionDenied,
			PermissionDenied,
			PermissionDeniedf,
			IsPermissionDenied,
		},
		{
			codes.ResourceExhausted,
			ResourceExhausted,
			ResourceExhaustedf,
			IsResourceExhausted,
		},
		{
			codes.Aborted,
			Aborted,
			Abortedf,
			IsAborted,
		},
		{
			codes.OutOfRange,
			OutOfRange,
			OutOfRangef,
			IsOutOfRange,
		},
		{
			codes.Unimplemented,
			Unimplemented,
			Unimplementedf,
			IsUnimplemeted,
		},
		{
			codes.Internal,
			Internal,
			Internalf,
			IsInternal,
		},
		{
			codes.Unavailable,
			Unavailable,
			Unavailablef,
			IsUnavailable,
		},
		{
			codes.DataLoss,
			DataLoss,
			DataLossf,
			IsDataLoss,
		},
		{
			codes.Unauthenticated,
			Unauthenticated,
			Unauthenticatedf,
			IsUnauthenticated,
		},
	}

	for _, tc := range tests {
		t.Log(tc, "config:")
		err := tc.ctor(nil)
		t.Log(err)
		if err.Code != tc.code {
			t.Fatal()
		}

		errf := tc.ctorf(nil, "hello:%s%s", "world", "!")
		t.Log(errf)
		if err.Error() == errf.Error() {
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

		we := Wrap(err)
		if Code(we) != Code(err) {
			t.Fatal(Code(we), Code(err))
		}

		wef := Wrapf(err, "text")
		if Code(wef) != Code(err) {
			t.Fatal(Code(wef), Code(err))
		}
	}
}

func TestIsCodeSuccessShortcut(t *testing.T) {
	ok := isCode(codes.OK, nil)
	if !ok {
		t.Fatal(ok)
	}
}
