// Code generated by "stringer -type=Code"; DO NOT EDIT.

package codes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-0]
	_ = x[Canceled-1]
	_ = x[Unknown-2]
	_ = x[InvalidArgument-3]
	_ = x[DeadlineExceeded-4]
	_ = x[NotFound-5]
	_ = x[AlreadyExists-6]
	_ = x[PermissionDenied-7]
	_ = x[ResourceExhausted-8]
	_ = x[FailedPrecondition-9]
	_ = x[Aborted-10]
	_ = x[OutOfRange-11]
	_ = x[Unimplemented-12]
	_ = x[Internal-13]
	_ = x[Unavailable-14]
	_ = x[DataLoss-15]
	_ = x[Unauthenticated-16]
}

const _Code_name = "OKCanceledUnknownInvalidArgumentDeadlineExceededNotFoundAlreadyExistsPermissionDeniedResourceExhaustedFailedPreconditionAbortedOutOfRangeUnimplementedInternalUnavailableDataLossUnauthenticated"

var _Code_index = [...]uint8{0, 2, 10, 17, 32, 48, 56, 69, 85, 102, 120, 127, 137, 150, 158, 169, 177, 192}

func (i Code) String() string {
	if i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
