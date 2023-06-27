package tErr

import "errors"

func newErr(str string) *Error {
	return &Error{
		err: errors.New(str),
	}
}

func newErrWithError(err error) *Error {
	return &Error{
		err: err,
	}
}
