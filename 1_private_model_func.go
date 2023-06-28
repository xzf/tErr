package tErr

import (
	"github.com/xzf/runtimeX"
	"strings"
)

func (t *tErr) Error() string {
	return t.msg
}

func (t *tErr) Track() Error {
	t.traceInfo = append(t.traceInfo, runtimeX.CallerLast())
	return t
}

func (t *tErr) TrackInfo() string {
	return t.msg + "\n" + strings.Join(t.traceInfo, "\n")
}

func (t *tErr) TrackSlice() []string {
	return t.traceInfo
}
