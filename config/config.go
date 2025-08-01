package config

import (
	"os"

	"github.com/joho/godotenv"
)

var GoogleAPIKey string

func init() {
	if err := godotenv.Load(); err != nil {
		// Don't panic in test environment
		if os.Getenv("TESTING") != "true" {
			panic("Error loading .env file")
		}
	}
	GoogleAPIKey = os.Getenv("GOOGLE_API_KEY")
}

const (
	ArticleTTSFile                = "article.txt"
	ArticleTTSSpeechRate          = 1.0
	SpellingPracticeFile          = "word_list.txt"
	SpellingPracticeWords         = 5
	SpellingPracticePauseDuration = 3 // seconds
)
