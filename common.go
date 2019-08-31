package gopherrs

import (
	"bytes"
	"fmt"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCError is a Go error that holds a grpc error code.
type GRPCError struct {
	Cause      error
	Code       codes.Code
	StackTrace []byte
	DescFmt    string
	DescData   []interface{}
}

// Error implements the error interface.
func (g *GRPCError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString("code:")
	buf.WriteString(g.Code.String())

	if len(g.DescFmt) > 0 {
		buf.WriteByte('\n')
		buf.WriteString("desc:")
		fmt.Fprintf(buf, g.DescFmt, g.DescData...)
	}

	if len(g.StackTrace) > 0 {
		buf.WriteByte('\n')
		buf.Write(g.StackTrace)
	}

	if g.Cause != nil {
		buf.WriteString("\n--------------------------------------------------------------------------------")
		buf.WriteString("\nCause:")
		buf.WriteString(g.Cause.Error())
	}

	return buf.String()
}

func newError(cause error, code codes.Code, format string, data []interface{}) *GRPCError {
	err := &GRPCError{cause, code, nil, "", nil}

	err.StackTrace = debug.Stack()
	err.DescFmt = format
	err.DescData = data
	return err
}

func isCode(code codes.Code, err error) bool {
	if code == codes.OK && err == nil {
		// early return if this is not an error...
		return true
	}

	if grpcErr, ok := err.(*GRPCError); ok {
		return code == grpcErr.Code
	}

	// fallback; check against the grpc status.Code package
	grpcCode := status.Code(err)
	return code == grpcCode
}

func Code(err error) codes.Code {
	if grpcErr, ok := err.(*GRPCError); ok {
		return grpcErr.Code
	}

	// fallback; check against the grpc status.Code package
	return status.Code(err)
}

func Canceled(cause error) *GRPCError {
	return Canceledf(cause, "")
}

func Canceledf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Canceled, f, data)
}

func IsCanceled(err error) bool {
	return isCode(codes.Canceled, err)
}

func Unknown(cause error) *GRPCError {
	return Unknownf(cause, "")
}

func Unknownf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Unknown, f, data)
}

func IsUnknown(err error) bool {
	return isCode(codes.Unknown, err)
}

func InvalidArgument(cause error) *GRPCError {
	return InvalidArgumentf(cause, "")
}

func InvalidArgumentf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.InvalidArgument, f, data)
}

func IsInvalidArgument(err error) bool {
	return isCode(codes.InvalidArgument, err)
}

func DeadlineExceeded(cause error) *GRPCError {
	return DeadlineExceededf(cause, "")
}

func DeadlineExceededf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.DeadlineExceeded, f, data)
}

func IsDeadlineExceeded(err error) bool {
	return isCode(codes.DeadlineExceeded, err)
}

func NotFound(cause error) *GRPCError {
	return NotFoundf(cause, "")
}

func NotFoundf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.NotFound, f, data)
}

func IsNotFound(err error) bool {
	return isCode(codes.NotFound, err)
}

func AlreadyExists(cause error) *GRPCError {
	return AlreadyExistsf(cause, "")
}

func AlreadyExistsf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.AlreadyExists, f, data)
}

func IsAlreadyExists(err error) bool {
	return isCode(codes.AlreadyExists, err)
}

func PermissionDenied(cause error) *GRPCError {
	return PermissionDeniedf(cause, "")
}

func PermissionDeniedf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.PermissionDenied, f, data)
}

func IsPermissionDenied(err error) bool {
	return isCode(codes.PermissionDenied, err)
}

func ResourceExhausted(cause error) *GRPCError {
	return ResourceExhaustedf(cause, "")
}

func ResourceExhaustedf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.ResourceExhausted, f, data)
}

func IsResourceExhausted(err error) bool {
	return isCode(codes.ResourceExhausted, err)
}

func FailedPrecondition(cause error) *GRPCError {
	return FailedPreconditionf(cause, "")
}

func FailedPreconditionf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.FailedPrecondition, f, data)
}

func IsFailedPrecondition(err error) bool {
	return isCode(codes.FailedPrecondition, err)
}

func Aborted(cause error) *GRPCError {
	return Abortedf(cause, "")
}

func Abortedf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Aborted, f, data)
}

func IsAborted(err error) bool {
	return isCode(codes.Aborted, err)
}

func OutOfRange(cause error) *GRPCError {
	return OutOfRangef(cause, "")
}

func OutOfRangef(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.OutOfRange, f, data)
}

func IsOutOfRange(err error) bool {
	return isCode(codes.OutOfRange, err)
}

func Unimplemented(cause error) *GRPCError {
	return Unimplementedf(cause, "")
}

func Unimplementedf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Unimplemented, f, data)
}

func IsUnimplemeted(err error) bool {
	return isCode(codes.Unimplemented, err)
}

func Internal(cause error) *GRPCError {
	return Internalf(cause, "")
}

func Internalf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Internal, f, data)
}

func IsInternal(err error) bool {
	return isCode(codes.Internal, err)
}

func Unavailable(cause error) *GRPCError {
	return Unavailablef(cause, "")
}

func Unavailablef(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Unavailable, f, data)
}

func IsUnavailable(err error) bool {
	return isCode(codes.Unavailable, err)
}

func DataLoss(cause error) *GRPCError {
	return DataLossf(cause, "")
}

func DataLossf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.DataLoss, f, data)
}

func IsDataLoss(err error) bool {
	return isCode(codes.DataLoss, err)
}

func Unauthenticated(cause error) *GRPCError {
	return Unauthenticatedf(cause, "")
}

func Unauthenticatedf(cause error, f string, data ...interface{}) *GRPCError {
	return newError(cause, codes.Unauthenticated, f, data)
}

func IsUnauthenticated(err error) bool {
	return isCode(codes.Unauthenticated, err)
}
