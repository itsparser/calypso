package errors

type BaseError struct {
	msg    string // errors description
	Offset int64  // where the errors occurred
	Code   int
}

func (e *BaseError) Error() string { return e.msg }
