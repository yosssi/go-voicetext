package voicetext

import (
	"encoding/json"
	"net/http"
	"testing"
)

func Test_newResult(t *testing.T) {
	var errMsg ttsErrMsg

	if err := json.Unmarshal([]byte(`{"error":{"message": "test msg"}}`), &errMsg); err != nil {
		t.Error(err)
	}

	result := newResult(http.StatusOK, []byte("test sound"), &errMsg)

	if result.Status != http.StatusOK {
		t.Errorf("result.Status should be %d [actual: %d]", http.StatusOK, result.Status)
	}

	if string(result.Sound) != "test sound" {
		t.Errorf("result.Sound should be %+v [actual: %+v]", []byte("test sound"), result.Sound)
	}

	if result.ErrMsg.String() != "test msg" {
		t.Errorf("result.ErrMsg.String() should be %s [actual: %s]", "test msg", result.ErrMsg.String())
	}
}
