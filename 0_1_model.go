package tErr

type Error interface {
	Error() string
	Track() Error
	TrackInfo() string
}
