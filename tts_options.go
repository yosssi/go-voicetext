package voicetext

// Emotions
const (
	EmotionHappiness = "happiness"
	EmotionAnger     = "anger"
	EmotionSadness   = "sadness"
)

// Defaults
const (
	defaultPitch  = 100
	defaultSpeed  = 100
	defaultVolume = 100
)

// Speakers
var (
	SpeakerShow = speaker{
		name:          "show",
		canUseEmotion: false,
	}
	SpeakerHaruka = speaker{
		name:          "haruka",
		canUseEmotion: true,
	}
	SpeakerHikari = speaker{
		name:          "hikari",
		canUseEmotion: true,
	}
	SpeakerTakeru = speaker{
		name:          "takeru",
		canUseEmotion: true,
	}

	emptySpeaker   = speaker{}
	defaultSpeaker = SpeakerShow
)

// TTSOptions represents options for the TTS method.
type TTSOptions struct {
	// Speaker represents a speaker of the speech.
	Speaker speaker
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
	if opts.Speaker == emptySpeaker {
		opts.Speaker = defaultSpeaker
	}
	if !opts.Speaker.canUseEmotion {
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
