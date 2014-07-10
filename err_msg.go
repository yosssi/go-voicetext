package voicetext

// ErrMsg is an interface for printing an error message
// returned by the API.
type ErrMsg interface {
	String() string
}
