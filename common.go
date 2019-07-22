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
func (g GRPCError) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString(g.Code.String())
	if len(g.StackTrace) > 0 {
		buf.WriteByte('\n')
		buf.WriteString(g.StackTrace)
	}

	return buf.String()
}

func newError(code codes.Code, opts []Option) *GRPCError {
	cfg := Configuration
	for _, opt := range opts {
		cfg = opt(cfg)
	}

	err := &GRPCError{code, ""}
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

func Canceled(opts ...Option) *GRPCError {
	return newError(codes.Canceled, opts)
}

func IsCanceled(err error) bool {
	return isCode(codes.Canceled, err)
}

func Unknown(opts ...Option) *GRPCError {
	return newError(codes.Unknown, opts)
}

func IsUnknown(err error) bool {
	return isCode(codes.Unknown, err)
}

func InvalidArgument(opts ...Option) *GRPCError {
	return newError(codes.InvalidArgument, opts)
}

func IsInvalidArgument(err error) bool {
	return isCode(codes.InvalidArgument, err)
}

func DeadlineExceeded(opts ...Option) *GRPCError {
	return newError(codes.DeadlineExceeded, opts)
}

func IsDeadlineExceeded(err error) bool {
	return isCode(codes.DeadlineExceeded, err)
}

func NotFound(opts ...Option) *GRPCError {
	return newError(codes.NotFound, opts)
}

func IsNotFound(err error) bool {
	return isCode(codes.NotFound, err)
}

func AlreadyExists(opts ...Option) *GRPCError {
	return newError(codes.AlreadyExists, opts)
}

func IsAlreadyExists(err error) bool {
	return isCode(codes.AlreadyExists, err)
}

func PermissionDenied(opts ...Option) *GRPCError {
	return newError(codes.PermissionDenied, opts)
}

func IsPermissionDenied(err error) bool {
	return isCode(codes.PermissionDenied, err)
}

func ResourceExhausted(opts ...Option) *GRPCError {
	return newError(codes.ResourceExhausted, opts)
}

func IsResourceExhausted(err error) bool {
	return isCode(codes.ResourceExhausted, err)
}

func FailedPrecondition(opts ...Option) *GRPCError {
	return newError(codes.FailedPrecondition, opts)
}

func IsFailedPrecondition(err error) bool {
	return isCode(codes.FailedPrecondition, err)
}

func Aborted(opts ...Option) *GRPCError {
	return newError(codes.Aborted, opts)
}

func IsAborted(err error) bool {
	return isCode(codes.Aborted, err)
}

func OutOfRange(opts ...Option) *GRPCError {
	return newError(codes.OutOfRange, opts)
}

func IsOutOfRange(err error) bool {
	return isCode(codes.OutOfRange, err)
}

func Unimplemented(opts ...Option) *GRPCError {
	return newError(codes.Unimplemented, opts)
}

func IsUnimplemeted(err error) bool {
	return isCode(codes.Unimplemented, err)
}

func Internal(opts ...Option) *GRPCError {
	return newError(codes.Internal, opts)
}

func IsInternal(err error) bool {
	return isCode(codes.Internal, err)
}

func Unavailable(opts ...Option) *GRPCError {
	return newError(codes.Unavailable, opts)
}

func IsUnavailable(err error) bool {
	return isCode(codes.Unavailable, err)
}

func DataLoss(opts ...Option) *GRPCError {
	return newError(codes.DataLoss, opts)
}

func IsDataLoss(err error) bool {
	return isCode(codes.DataLoss, err)
}

func Unauthenticated(opts ...Option) *GRPCError {
	return newError(codes.Unauthenticated, opts)
}

func IsUnauthenticated(err error) bool {
	return isCode(codes.Unauthenticated, err)
}
