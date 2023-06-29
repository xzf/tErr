package tErr

import (
	"github.com/xzf/runtimeX"
	"strings"
)

func (t *tErr) Error() string {
	return t.msg
}

func (t *tErr) Trace() Error {
	t.traceInfo = append(t.traceInfo, runtimeX.CallerLast())
	return t
}

func (t *tErr) TraceInfo() string {
	return t.msg + "\n" + strings.Join(t.traceInfo, "\n")
}

func (t *tErr) TraceSlice() []string {
	return t.traceInfo
}
