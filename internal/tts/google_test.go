package tts

import (
	"testing"
)

func TestNewGoogleTTS(t *testing.T) {
	apiKey := "test-api-key"
	tts := NewGoogleTTS(apiKey)

	if tts.APIKey != apiKey {
		t.Errorf("Expected API key %q, got %q", apiKey, tts.APIKey)
	}

	if tts.Client == nil {
		t.Error("Expected HTTP client to be initialized")
	}
}

func TestGetRandomVoiceName(t *testing.T) {
	tts := NewGoogleTTS("test-key")

	// Test multiple calls to ensure we get different voices
	voices := make(map[string]bool)
	for i := 0; i < 10; i++ {
		voice := tts.GetRandomVoiceName()
		if voice == "" {
			t.Error("Expected non-empty voice name")
		}
		voices[voice] = true
	}

	// With 10 calls, we should get at least 2 different voices
	if len(voices) < 2 {
		t.Errorf("Expected variety in voice selection, got only %d unique voices", len(voices))
	}
}
