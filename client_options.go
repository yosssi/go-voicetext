package voicetext

// Defaults
const (
	defaultVersion = "v1"
)

// ClientOptions represents options for a client.
type ClientOptions struct {
	// Version represents the version of VoiceText Web API.
	Version string
}

// setDefaults sets defaults to the client options.
func (opts *ClientOptions) setDefaults() {
	if opts.Version == "" {
		opts.Version = defaultVersion
	}
}
