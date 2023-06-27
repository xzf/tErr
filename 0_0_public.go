package tErr

func New(str string) *Error {
	return newErr(str)
}

func Track(err error) *Error {
	return newErrWithError(err)
}
