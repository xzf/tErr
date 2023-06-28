package tErr

import "github.com/xzf/runtimeX"

func New(str string) Error {
	return &tErr{
		msg: str,
		traceInfo: []string{
			runtimeX.CallerLast(),
		},
	}
}

func NewByError(err error) Error {
	return &tErr{
		msg: err.Error(),
		traceInfo: []string{
			runtimeX.CallerLast(),
		},
	}
}
