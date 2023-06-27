package tErr

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Trace() Error {
	return e
}

func (e Error) GoError() error {
	return e.err
}

func (e Error) TraceInfo() string {
	return e.err.Error()
}
