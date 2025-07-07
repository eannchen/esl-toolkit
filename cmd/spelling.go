package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"esl-toolkit/config"
	"esl-toolkit/internal/audio"
	"esl-toolkit/internal/tts"
	"esl-toolkit/internal/words"
)

// Spelling Practice Tool: Plays random words using TTS for spelling practice.
func RunSpellingPractice() {
	words, err := words.LoadWords(config.SpellingPracticeFile)
	if err != nil {
		fmt.Println("Error loading words:", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	googleTTS := tts.NewGoogleTTS(config.GoogleAPIKey)

	fmt.Println("Total: ", len(words))

	for i := 0; i < config.SpellingPracticeWords; i++ {
		wordIndex := rand.Intn(len(words))
		word := words[wordIndex]
		voiceName := googleTTS.GetRandomVoiceName()
		audioData, err := googleTTS.SynthesizeSpeech(word, voiceName)
		if err != nil {
			fmt.Println("Error getting audio data:", err)
			continue
		}
		for j := 0; j < 2; j++ {
			if err = audio.PlayAudio(audioData); err != nil {
				fmt.Println("Error playing audio:", err)
				continue
			}
			time.Sleep(time.Duration(config.SpellingPracticePauseDuration) * time.Second)
		}
		fmt.Println(i+1, word)
		words = append(words[:wordIndex], words[wordIndex+1:]...)
	}
}
