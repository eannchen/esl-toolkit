package tts

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var voiceNameList = []string{
	"en-GB-Neural2-A", "en-GB-Neural2-B", "en-GB-Neural2-C", "en-GB-Neural2-D", "en-GB-Neural2-F",
	"en-GB-News-G", "en-GB-News-H", "en-GB-News-I", "en-GB-News-J", "en-GB-News-K", "en-GB-News-L", "en-GB-News-M",
	"en-GB-Wavenet-A", "en-GB-Wavenet-B", "en-GB-Wavenet-C", "en-GB-Wavenet-D", "en-GB-Wavenet-F",
}

type AudioContent struct {
	AudioContent string `json:"audioContent"`
}

type GoogleTTS struct {
	APIKey string
	Client *http.Client
}

func NewGoogleTTS(apiKey string) *GoogleTTS {
	return &GoogleTTS{
		APIKey: apiKey,
		Client: &http.Client{},
	}
}

func (g *GoogleTTS) GetRandomVoiceName() string {
	if len(voiceNameList) == 0 {
		return ""
	}
	return voiceNameList[rand.Intn(len(voiceNameList))]
}

func (g *GoogleTTS) SynthesizeSpeech(text, voiceName string) ([]byte, error) {
	apiURL := fmt.Sprintf("https://texttospeech.googleapis.com/v1/text:synthesize?key=%s", g.APIKey)

	reqBody := fmt.Sprintf(`{"input": {"text": "%s"}, "voice": {"languageCode": "en-GB", "name": "%s"}, "audioConfig": {"audioEncoding": "mp3"}}`, text, voiceName)

	resp, err := g.Client.Post(apiURL, "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var ac AudioContent
	if err = json.Unmarshal(audioData, &ac); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return base64.StdEncoding.DecodeString(ac.AudioContent)
}

func (g *GoogleTTS) SynthesizeSpeechWithRate(text, voiceName string, speakingRate float64) ([]byte, error) {
	apiURL := fmt.Sprintf("https://texttospeech.googleapis.com/v1/text:synthesize?key=%s", g.APIKey)
	reqBody := fmt.Sprintf(
		`{"input": {"text": %q}, "voice": {"languageCode": "en-GB", "name": %q}, "audioConfig": {"audioEncoding": "mp3", "speakingRate": %.2f}}`,
		text, voiceName, speakingRate,
	)
	resp, err := g.Client.Post(apiURL, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var ac AudioContent
	if err = json.Unmarshal(audioData, &ac); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	if ac.AudioContent == "" {
		return nil, fmt.Errorf("audioContent field is empty in the API response")
	}
	return base64.StdEncoding.DecodeString(ac.AudioContent)
}
