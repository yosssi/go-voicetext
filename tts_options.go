package voicetext

// Speakers
const (
	SpeakerShow   = "show"
	SpeakerHaruka = "haruka"
	SpeakerHikari = "hikari"
	SpeakerTakeru = "takeru"
)

// Emotions
const (
	EmotionHappiness = "happiness"
	EmotionAnger     = "anger"
	EmotionSadness   = "sadness"
)

// Defaults
const (
	defaultSpeaker = SpeakerShow
	defaultPitch   = 100
	defaultSpeed   = 100
	defaultVolume  = 100
)

// TTSOptions represents options for the TTS method.
type TTSOptions struct {
	// Speaker represents a speaker of the speech.
	Speaker string
	// Emotion represents an emotion of the speech.
	Emotion string
	// EmotionLevel represents an emotion level of the speech.
	EmotionLevel string
	// Pitch represents a pitch of the speech.
	Pitch int
	// Speed represents a speed of the speech.
	Speed int
	// Volume represents a volume of the speech.
	Volume int
}

// setDefaults sets defaults to the TTS method's options.
func (opts *TTSOptions) setDefaults() {
	if opts.Speaker == "" {
		opts.Speaker = defaultSpeaker
	}
	if opts.Speaker == SpeakerShow {
		opts.Emotion = ""
	}
	if opts.Pitch == 0 {
		opts.Pitch = defaultPitch
	}
	if opts.Speed == 0 {
		opts.Speed = defaultSpeed
	}
	if opts.Volume == 0 {
		opts.Volume = defaultVolume
	}
}
