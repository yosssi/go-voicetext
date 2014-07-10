package voicetext

import (
	"encoding/json"
	"testing"
)

func Test_ttsErrMsg_String(t *testing.T) {
	var errMsg ttsErrMsg

	if err := json.Unmarshal([]byte(`{"error":{"message": "test msg"}}`), &errMsg); err != nil {
		t.Error(err)
	}

	if errMsg.String() != "test msg" {
		t.Errorf("ttsErrMsg.String() should be %s [actual: %s]", "test msg", errMsg.String())
	}
}
