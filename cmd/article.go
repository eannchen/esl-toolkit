package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"esl-toolkit/internal/audio"
	"esl-toolkit/internal/tts"
)

func RunArticleTTS(apiKey, articleFile string) {
	mp3FilePath := strings.TrimSuffix(articleFile, ".txt") + ".mp3"
	if _, err := os.Stat(mp3FilePath); err == nil {
		if err := audio.PlayAudioFromFile(mp3FilePath); err != nil {
			fmt.Println("Error playing audio:", err)
		}
		return
	}

	text, err := os.ReadFile(articleFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	voiceNames := []string{
		"en-GB-Neural2-B", "en-GB-Neural2-F", "en-GB-News-J",
		"en-GB-News-L", "en-GB-Wavenet-B", "en-GB-Wavenet-F",
	}
	rand.Seed(time.Now().UnixNano())
	voiceName := voiceNames[rand.Intn(len(voiceNames))]
	speakingRate := 0.85

	chunks := splitTextIntoChunks(string(text), 5000)
	var combinedAudio []byte
	googleTTS := tts.NewGoogleTTS(apiKey)

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
