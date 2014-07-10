package voicetext

import "testing"

func TestTTSOptions_setDefaults(t *testing.T) {
	opts := &TTSOptions{}
	opts.setDefaults()
	if opts.Speaker != defaultSpeaker {
		t.Errorf("opts.Speaker should be %s [actual: %s]", defaultSpeaker, opts.Speaker)
	}
	if opts.Emotion != "" {
		t.Errorf("opts.Speaker should be %s [actual: %s]", `""`, opts.Emotion)
	}
	if opts.Pitch != defaultPitch {
		t.Errorf("opts.Speaker should be %d [actual: %d]", defaultPitch, opts.Pitch)
	}
	if opts.Speed != defaultSpeed {
		t.Errorf("opts.Speed should be %s [actual: %s]", defaultSpeed, opts.Speed)
	}
	if opts.Volume != defaultVolume {
		t.Errorf("opts.Volume should be %s [actual: %s]", defaultVolume, opts.Volume)
	}
}
