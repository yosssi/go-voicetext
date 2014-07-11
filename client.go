package voicetext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// API URL and paths
const (
	apiURL     = "https://api.voicetext.jp/%s/%s"
	apiPathTTS = "tts"
)

// HTTP methods
const (
	httpPOST = "POST"
)

// Client is an interface for calling VoiceText Web API.
type Client interface {
	// TTS calls the tts API.
	TTS(text string, opts *TTSOptions) (*Result, error)
}

// client represens a client for VoiceText Web API.
type client struct {
	// apiKey represents an API key for VoiceText Web API.
	apiKey string
	// version represents the version of VoiceText Web API.
	version string
}

// TTS calls the tts API.
func (c *client) TTS(text string, opts *TTSOptions) (*Result, error) {
	if opts == nil {
		opts = &TTSOptions{}
	}

	opts.setDefaults()

	// Post the data
	res, err := c.post(ttsURL(c.version), ttsValues(text, opts))
	if err != nil {
		return nil, err
	}

	// Read the response
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	status := res.StatusCode

	// Return the error message if the HTTP status code does not equal to StatusOK.
	if status != http.StatusOK {
		var errMsg TTSErrMsg
		if err := json.Unmarshal(b, &errMsg); err != nil {
			result := newResult(
				status,
				nil,
				&TTSErrMsg{
					Err: TTSErrMsgError{
						Message: string(b),
					},
				},
			)
			return result, nil
		}
		return newResult(status, nil, &errMsg), nil
	}

	return newResult(status, b, nil), nil
}

// post posts data to the specified URL.
func (c *client) post(url, values string) (*http.Response, error) {
	req, err := http.NewRequest(httpPOST, url, bytes.NewBufferString(values))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(values)))

	return http.DefaultClient.Do(req)
}

// NewClient creates and
func NewClient(apiKey string, opts *ClientOptions) *client {
	if opts == nil {
		opts = &ClientOptions{}
	}

	opts.setDefaults()

	return &client{
		apiKey:  apiKey,
		version: opts.Version,
	}
}

// ttsURL creates and returns an URL of the TTS API.
func ttsURL(version string) string {
	return fmt.Sprintf(apiURL, version, apiPathTTS)
}

// ttsValues creates and returns encoded parameters for the TTS API.
func ttsValues(text string, opts *TTSOptions) string {
	values := url.Values{}
	values.Add("text", text)
	values.Add("speaker", opts.Speaker.name)
	values.Add("emotion", opts.Emotion)
	values.Add("emotion_level", opts.EmotionLevel)
	values.Add("pitch", strconv.Itoa(opts.Pitch))
	values.Add("speed", strconv.Itoa(opts.Speed))
	values.Add("volume", strconv.Itoa(opts.Volume))
	return values.Encode()
}
