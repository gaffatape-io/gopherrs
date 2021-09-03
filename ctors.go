package gopherrs

import (
	"github.com/gaffatape-io/gopherrs/codes"
)

func Canceled(msg string) *E {
	return NewE(codes.Canceled, msg)
}
func Canceledf(msg string, args ...interface{}) *E {
	return NewEf(codes.Canceled, msg, args...)
}
func WrapCanceled(cause error, msg string) *E {
	return Wrap(cause, codes.Canceled, msg)
}
func WrapCanceledf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Canceled, msg, args...)
}
func Unknown(msg string) *E {
	return NewE(codes.Unknown, msg)
}
func Unknownf(msg string, args ...interface{}) *E {
	return NewEf(codes.Unknown, msg, args...)
}
func WrapUnknown(cause error, msg string) *E {
	return Wrap(cause, codes.Unknown, msg)
}
func WrapUnknownf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Unknown, msg, args...)
}
func InvalidArgument(msg string) *E {
	return NewE(codes.InvalidArgument, msg)
}
func InvalidArgumentf(msg string, args ...interface{}) *E {
	return NewEf(codes.InvalidArgument, msg, args...)
}
func WrapInvalidArgument(cause error, msg string) *E {
	return Wrap(cause, codes.InvalidArgument, msg)
}
func WrapInvalidArgumentf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.InvalidArgument, msg, args...)
}
func DeadlineExceeded(msg string) *E {
	return NewE(codes.DeadlineExceeded, msg)
}
func DeadlineExceededf(msg string, args ...interface{}) *E {
	return NewEf(codes.DeadlineExceeded, msg, args...)
}
func WrapDeadlineExceeded(cause error, msg string) *E {
	return Wrap(cause, codes.DeadlineExceeded, msg)
}
func WrapDeadlineExceededf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.DeadlineExceeded, msg, args...)
}
func PermissionDenied(msg string) *E {
	return NewE(codes.PermissionDenied, msg)
}
func PermissionDeniedf(msg string, args ...interface{}) *E {
	return NewEf(codes.PermissionDenied, msg, args...)
}
func WrapPermissionDenied(cause error, msg string) *E {
	return Wrap(cause, codes.PermissionDenied, msg)
}
func WrapPermissionDeniedf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.PermissionDenied, msg, args...)
}
func ResourceExhausted(msg string) *E {
	return NewE(codes.ResourceExhausted, msg)
}
func ResourceExhaustedf(msg string, args ...interface{}) *E {
	return NewEf(codes.ResourceExhausted, msg, args...)
}
func WrapResourceExhausted(cause error, msg string) *E {
	return Wrap(cause, codes.ResourceExhausted, msg)
}
func WrapResourceExhaustedf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.ResourceExhausted, msg, args...)
}
func FailedPrecondition(msg string) *E {
	return NewE(codes.FailedPrecondition, msg)
}
func FailedPreconditionf(msg string, args ...interface{}) *E {
	return NewEf(codes.FailedPrecondition, msg, args...)
}
func WrapFailedPrecondition(cause error, msg string) *E {
	return Wrap(cause, codes.FailedPrecondition, msg)
}
func WrapFailedPreconditionf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.FailedPrecondition, msg, args...)
}
func Aborted(msg string) *E {
	return NewE(codes.Aborted, msg)
}
func Abortedf(msg string, args ...interface{}) *E {
	return NewEf(codes.Aborted, msg, args...)
}
func WrapAborted(cause error, msg string) *E {
	return Wrap(cause, codes.Aborted, msg)
}
func WrapAbortedf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Aborted, msg, args...)
}
func OutOfRange(msg string) *E {
	return NewE(codes.OutOfRange, msg)
}
func OutOfRangef(msg string, args ...interface{}) *E {
	return NewEf(codes.OutOfRange, msg, args...)
}
func WrapOutOfRange(cause error, msg string) *E {
	return Wrap(cause, codes.OutOfRange, msg)
}
func WrapOutOfRangef(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.OutOfRange, msg, args...)
}
func Unimplemented(msg string) *E {
	return NewE(codes.Unimplemented, msg)
}
func Unimplementedf(msg string, args ...interface{}) *E {
	return NewEf(codes.Unimplemented, msg, args...)
}
func WrapUnimplemented(cause error, msg string) *E {
	return Wrap(cause, codes.Unimplemented, msg)
}
func WrapUnimplementedf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Unimplemented, msg, args...)
}
func Internal(msg string) *E {
	return NewE(codes.Internal, msg)
}
func Internalf(msg string, args ...interface{}) *E {
	return NewEf(codes.Internal, msg, args...)
}
func WrapInternal(cause error, msg string) *E {
	return Wrap(cause, codes.Internal, msg)
}
func WrapInternalf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Internal, msg, args...)
}
func Unavailable(msg string) *E {
	return NewE(codes.Unavailable, msg)
}
func Unavailablef(msg string, args ...interface{}) *E {
	return NewEf(codes.Unavailable, msg, args...)
}
func WrapUnavailable(cause error, msg string) *E {
	return Wrap(cause, codes.Unavailable, msg)
}
func WrapUnavailablef(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Unavailable, msg, args...)
}
func DataLoss(msg string) *E {
	return NewE(codes.DataLoss, msg)
}
func DataLossf(msg string, args ...interface{}) *E {
	return NewEf(codes.DataLoss, msg, args...)
}
func WrapDataLoss(cause error, msg string) *E {
	return Wrap(cause, codes.DataLoss, msg)
}
func WrapDataLossf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.DataLoss, msg, args...)
}
func Unauthenticated(msg string) *E {
	return NewE(codes.Unauthenticated, msg)
}
func Unauthenticatedf(msg string, args ...interface{}) *E {
	return NewEf(codes.Unauthenticated, msg, args...)
}
func WrapUnauthenticated(cause error, msg string) *E {
	return Wrap(cause, codes.Unauthenticated, msg)
}
func WrapUnauthenticatedf(cause error, msg string, args ...interface{}) *E {
	return Wrapf(cause, codes.Unauthenticated, msg, args...)
}
