package voicetext

// TTSErrMsg represents an error meesage of the TTS method.
type ttsErrMsg struct {
	Error struct {
		Meesage string `json:"message"`
	} `json:"error"`
}

// String returns the error message.
func (msg *ttsErrMsg) String() string {
	return msg.Error.Meesage
}
