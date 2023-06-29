package tErr

type Error interface {
	Error() string
	Trace() Error
	TraceInfo() string
	TraceSlice() []string
}
