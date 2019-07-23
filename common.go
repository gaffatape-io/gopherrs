package gopherrs

import (
	"bytes"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Configuration holds the global configuration used when creating errors.
	Configuration Config = Config{
		GenerateStacktrace: true,
	}
)

// GRPCError is a Go error that holds a grpc error code.
type GRPCError struct {
	Cause      error
	Code       codes.Code
	StackTrace string
}

// Config is used to configure the error created.
type Config struct {
	GenerateStacktrace bool
}

// Option is used to override the global config for individual errors.
type Option func(cfg Config) Config

// NoStackTrace disables stack trace generation for this error.
func NoStackTrace(cfg Config) Config {
	cfg.GenerateStacktrace = false
	return cfg
}

// Error implements the error interface.
func (g *GRPCError) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString(g.Code.String())
	if len(g.StackTrace) > 0 {
		buf.WriteByte('\n')
		buf.WriteString(g.StackTrace)
	}

	if g.Cause != nil {
		buf.WriteString("\n--------------------------------------------------------------------------------")
		buf.WriteString("\nCause:")
		buf.WriteString(g.Cause.Error())
	}

	return buf.String()
}

func newError(cause error, code codes.Code, opts []Option) *GRPCError {
	cfg := Configuration
	for _, opt := range opts {
		cfg = opt(cfg)
	}

	err := &GRPCError{cause, code, ""}
	if cfg.GenerateStacktrace {
		err.StackTrace = string(debug.Stack())
	}
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

func Canceled(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Canceled, opts)
}

func IsCanceled(err error) bool {
	return isCode(codes.Canceled, err)
}

func Unknown(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Unknown, opts)
}

func IsUnknown(err error) bool {
	return isCode(codes.Unknown, err)
}

func InvalidArgument(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.InvalidArgument, opts)
}

func IsInvalidArgument(err error) bool {
	return isCode(codes.InvalidArgument, err)
}

func DeadlineExceeded(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.DeadlineExceeded, opts)
}

func IsDeadlineExceeded(err error) bool {
	return isCode(codes.DeadlineExceeded, err)
}

func NotFound(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.NotFound, opts)
}

func IsNotFound(err error) bool {
	return isCode(codes.NotFound, err)
}

func AlreadyExists(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.AlreadyExists, opts)
}

func IsAlreadyExists(err error) bool {
	return isCode(codes.AlreadyExists, err)
}

func PermissionDenied(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.PermissionDenied, opts)
}

func IsPermissionDenied(err error) bool {
	return isCode(codes.PermissionDenied, err)
}

func ResourceExhausted(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.ResourceExhausted, opts)
}

func IsResourceExhausted(err error) bool {
	return isCode(codes.ResourceExhausted, err)
}

func FailedPrecondition(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.FailedPrecondition, opts)
}

func IsFailedPrecondition(err error) bool {
	return isCode(codes.FailedPrecondition, err)
}

func Aborted(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Aborted, opts)
}

func IsAborted(err error) bool {
	return isCode(codes.Aborted, err)
}

func OutOfRange(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.OutOfRange, opts)
}

func IsOutOfRange(err error) bool {
	return isCode(codes.OutOfRange, err)
}

func Unimplemented(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Unimplemented, opts)
}

func IsUnimplemeted(err error) bool {
	return isCode(codes.Unimplemented, err)
}

func Internal(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Internal, opts)
}

func IsInternal(err error) bool {
	return isCode(codes.Internal, err)
}

func Unavailable(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Unavailable, opts)
}

func IsUnavailable(err error) bool {
	return isCode(codes.Unavailable, err)
}

func DataLoss(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.DataLoss, opts)
}

func IsDataLoss(err error) bool {
	return isCode(codes.DataLoss, err)
}

func Unauthenticated(cause error, opts ...Option) *GRPCError {
	return newError(cause, codes.Unauthenticated, opts)
}

func IsUnauthenticated(err error) bool {
	return isCode(codes.Unauthenticated, err)
}
