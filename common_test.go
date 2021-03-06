package gopherrs

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"fmt"
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

		werr := Wrap(err)
		if Code(werr) != Code(err) {
			t.Fatal(Code(err), Code(err))
		}

		werrf := Wrapf(err, "wrapper")
		if Code(werrf) != Code(err) {
			t.Fatal(Code(err), Code(err))
		}
	}
}

func TestIsCodeSuccessShortcut(t *testing.T) {
	ok := isCode(codes.OK, nil)
	if !ok {
		t.Fatal(ok)
	}
}

func TestWrap(t *testing.T) {
	err := fmt.Errorf(t.Name())	
	WrapError(&err)

	if _, ok := err.(*GRPCError); !ok {
		t.Fatal("Wrap failed", err)
	}

	nilErrorMaker := func() error {
		t.Log("no error")
		return nil
	}

	err2 := nilErrorMaker()
	if err2 != nil {
		t.Fatal(err2)
	}
	WrapError(&err2)
	if err2 != nil {
		t.Fatal(err2)
	}	
}
