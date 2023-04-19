package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type AudioContent struct {
	AudioContent string `json:"audioContent"`
}

func main() {
	// Replace "YOUR_API_KEY" with your Google Text-to-Speech API key
	apiKey := ""

	// Read the word list from a file
	file, err := os.Open("word_list.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read each line (word) from the file and store them in a slice
	words := make([]string, 0)
	var word string
	for {
		_, err := fmt.Fscanf(file, "%s\n", &word)
		if err != nil {
			break
		}
		words = append(words, word)
	}

	// Set the number of words to play
	numWords := 10

	// Set the pause duration between each word (in seconds)
	pauseDuration := 6

	// Shuffle the words slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	// Create a HTTP client with timeout
	client := &http.Client{Timeout: time.Second * 10}

	time.Sleep(2 * time.Second)

	fmt.Println(len(words))

	// Loop through the words slice and play each word twice
	for i := 0; i < numWords; i++ {
		wordIndex := rand.Intn(len(words))
		word := words[wordIndex]

		// Build the Google Text-to-Speech API URL
		apiURL := fmt.Sprintf("https://texttospeech.googleapis.com/v1/text:synthesize?key=%s", apiKey)

		// Random name
		voiceNameList := []string{
			"en-GB-Neural2-A",
			"en-GB-Neural2-B",
			"en-GB-Neural2-C",
			"en-GB-Neural2-D",
			"en-GB-Neural2-F",
			"en-GB-News-G",
			"en-GB-News-H",
			"en-GB-News-I",
			"en-GB-News-J",
			"en-GB-News-K",
			"en-GB-News-L",
			"en-GB-News-M",
			"en-GB-Wavenet-A",
			"en-GB-Wavenet-B",
			"en-GB-Wavenet-C",
			"en-GB-Wavenet-D",
			"en-GB-Wavenet-F",
		}
		rand.Seed(time.Now().UnixNano())
		voiceName := voiceNameList[rand.Intn(len(voiceNameList))]

		// Build the request body with the word
		reqBody := fmt.Sprintf(`{"input": {"text": "%s"}, "voice": {"languageCode": "en-GB", "name": "%s"}, "audioConfig": {"audioEncoding": "mp3"}}`, word, voiceName)

		// Send the request to the Google Text-to-Speech API
		resp, err := client.Post(apiURL, "application/json", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		// Read the response body into a byte slice
		audioData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		var ac AudioContent
		if err = json.Unmarshal(audioData, &ac); err != nil {
			fmt.Println("Error decoding JSON: ", err)
			return
		}
		decodeAC, err := base64.StdEncoding.DecodeString(ac.AudioContent)
		if err != nil {
			fmt.Println("Error decoding JSON: ", err)
			return
		}

		// Play the audio file twice with a pause between each play
		for j := 0; j < 2; j++ {
			if err = playAudio([]byte(decodeAC)); err != nil {
				fmt.Println("Error play audio: ", err)
				continue
			}
			time.Sleep(time.Duration(pauseDuration) * time.Second)
		}

		// Answer
		fmt.Println(i+1, word)

		// Remove the played word from the words slice
		words = append(words[:wordIndex], words[wordIndex+1:]...)
	}
}

func playAudio(audioData []byte) error {
	tmpfile, err := ioutil.TempFile("", "audio*.mp3")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	// Write the audio data to the temporary file
	if _, err := tmpfile.Write(audioData); err != nil {
		return err
	}

	// Close the file to ensure that all data is written to disk
	if err := tmpfile.Close(); err != nil {
		return err
	}

	// Play the audio file using the 'afplay' command
	cmd := exec.Command("afplay", tmpfile.Name())
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
