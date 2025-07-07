package cmd

import (
	"fmt"
	"os"
	"strings"

	"esl-toolkit/config"
	"esl-toolkit/internal/audio"
	"esl-toolkit/internal/tts"
)

// Article TTS Tool: Reads an article aloud using TTS and saves/plays the audio.
func RunArticleTTS() {
	mp3FilePath := strings.TrimSuffix(config.ArticleTTSFile, ".txt") + ".mp3"
	if _, err := os.Stat(mp3FilePath); err == nil {
		if err := audio.PlayAudioFromFile(mp3FilePath); err != nil {
			fmt.Println("Error playing audio:", err)
		}
		return
	}

	text, err := os.ReadFile(config.ArticleTTSFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	speakingRate := 0.85

	chunks := splitTextIntoChunks(string(text), 5000)
	var combinedAudio []byte

	googleTTS := tts.NewGoogleTTS(config.GoogleAPIKey)
	voiceName := googleTTS.GetRandomVoiceName()

	for _, chunk := range chunks {
		audioData, err := googleTTS.SynthesizeSpeechWithRate(chunk, voiceName, speakingRate)
		if err != nil {
			fmt.Println("Error synthesizing chunk:", err)
			return
		}
		combinedAudio = append(combinedAudio, audioData...)
	}

	if err := os.WriteFile(mp3FilePath, combinedAudio, 0644); err != nil {
		fmt.Println("Error saving audio:", err)
		return
	}
	if err := audio.PlayAudioFromFile(mp3FilePath); err != nil {
		fmt.Println("Error playing audio:", err)
	}
}

func splitTextIntoChunks(text string, maxSize int) []string {
	var chunks []string
	for len(text) > maxSize {
		idx := strings.LastIndex(text[:maxSize], " ")
		if idx == -1 {
			idx = maxSize
		}
		chunks = append(chunks, text[:idx])
		text = text[idx:]
	}
	chunks = append(chunks, text)
	return chunks
}
