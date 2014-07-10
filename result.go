package voicetext

// Result represents a result of the API call.
type Result struct {
	// Status represents an HTTP status.
	Status int
	// Sound represents a sound data.
	Sound []byte
	// ErrMsg represents an error message.
	ErrMsg ErrMsg
}

// newResult creates a result of the API call.
func newResult(status int, sound []byte, errMsg ErrMsg) *Result {
	return &Result{
		Status: status,
		Sound:  sound,
		ErrMsg: errMsg,
	}
}
