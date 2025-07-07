package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"els-spelling/internal/audio"
	"els-spelling/internal/tts"
	"els-spelling/internal/words"
)

const (
	apiKey        = "" // Set your Google Text-to-Speech API key here
	wordListFile  = "word_list.txt"
	numWords      = 5
	pauseDuration = 3 // seconds
)

func RunSpellingPractice() {

	words, err := words.LoadWords(wordListFile)
	if err != nil {
		fmt.Println("Error loading words:", err)
		os.Exit(1)
	}

	// Shuffle the words slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	// Initialize the Google Text-to-Speech API with the provided API key
	googleTTS := tts.NewGoogleTTS(apiKey)

	fmt.Println("Total: ", len(words))

	// Loop through the words slice and play each word
	for i := 0; i < numWords; i++ {
		wordIndex := rand.Intn(len(words))
		word := words[wordIndex]

		// Get a random voice name
		voiceName := googleTTS.GetRandomVoiceName()

		// Send the request to the Google Text-to-Speech API and get audio data
		audioData, err := googleTTS.SynthesizeSpeech(word, voiceName)
		if err != nil {
			fmt.Println("Error getting audio data:", err)
			continue
		}

		// Play the audio file twice with a pause between each play
		for j := 0; j < 2; j++ {
			if err = audio.PlayAudio(audioData); err != nil {
				fmt.Println("Error playing audio:", err)
				continue
			}
			time.Sleep(time.Duration(pauseDuration) * time.Second)
		}

		// Output the played word
		fmt.Println(i+1, word)

		// Remove the played word from the words slice
		words = append(words[:wordIndex], words[wordIndex+1:]...)
	}
}
