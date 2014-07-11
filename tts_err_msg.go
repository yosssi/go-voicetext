package voicetext

// TTSErrMsg represents an error meesage of the TTS method.
type TTSErrMsg struct {
	Err TTSErrMsgError `json:"error"`
}

// TTSErrMsgError represents an error of the tts error message.
type TTSErrMsgError struct {
	Message string `json:"message"`
}

// String returns the error message.
func (msg *TTSErrMsg) String() string {
	return msg.Err.Message
}
