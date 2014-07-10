package voicetext

import "testing"

func TestClientOptions_setDefaults(t *testing.T) {
	opts := &ClientOptions{}
	opts.setDefaults()
	actual := opts.Version
	expected := defaultVersion
	if actual != expected {
		t.Errorf("opts.Version should be %s [actual: %s]", expected, actual)
	}
}
