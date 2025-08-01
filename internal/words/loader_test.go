package words

import (
	"os"
	"testing"
)

func TestLoadWords(t *testing.T) {
	// Create a temporary test file
	content := "apple\nbeautiful\ncomputer\ndictionary\nelephant"
	tmpfile, err := os.CreateTemp("", "test_words_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test content
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test loading words
	words, err := LoadWords(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadWords failed: %v", err)
	}

	expected := []string{"apple", "beautiful", "computer", "dictionary", "elephant"}
	if len(words) != len(expected) {
		t.Errorf("Expected %d words, got %d", len(expected), len(words))
	}

	for i, word := range expected {
		if words[i] != word {
			t.Errorf("Expected word %d to be %q, got %q", i, word, words[i])
		}
	}
}

func TestLoadWordsEmptyFile(t *testing.T) {
	// Create empty test file
	tmpfile, err := os.CreateTemp("", "test_empty_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	words, err := LoadWords(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadWords failed with empty file: %v", err)
	}

	if len(words) != 0 {
		t.Errorf("Expected 0 words from empty file, got %d", len(words))
	}
}

func TestLoadWordsNonExistentFile(t *testing.T) {
	_, err := LoadWords("nonexistent_file.txt")
	if err == nil {
		t.Error("Expected error when loading non-existent file")
	}
}
