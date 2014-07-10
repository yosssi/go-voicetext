package voicetext

import (
	"os"
	"testing"
)

func Test_client_TTS_normalEnd(t *testing.T) {
	client := NewClient(os.Getenv("VOICETEXT_API_KEY"), nil)
	_, err := client.TTS("hello", nil)
	if err != nil {
		t.Error(err)
	}
}

func Test_client_TTS_errMsgReturned(t *testing.T) {
	client := NewClient("", nil)
	_, err := client.TTS("hello", nil)
	if err != nil {
		t.Error(err)
	}
}

func Test_client_TTS_jsonUnmarshalReturnsError(t *testing.T) {
	client := NewClient("", &ClientOptions{
		Version: "notExistVersion",
	})
	_, err := client.TTS("hello", nil)
	if err == nil {
		t.Error("an error should be returned.")
	}
}
