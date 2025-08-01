# ESL Toolkit

A collection of practical tools for English as a Second Language (ESL) learners. This toolkit provides Text-to-Speech (TTS) functionality for spelling practice and article reading, making it easier to improve English pronunciation and listening skills.

## Features

- **📝 Spelling Practice**: Randomly plays words using Google TTS for spelling practice sessions
- **📖 Article TTS**: Converts text articles to speech and plays them aloud
- **⚡ MP3 Processing**: Batch processes MP3 files to adjust playback speed using FFmpeg

## Quick Start

### Prerequisites

- **Go 1.18+** - [Download here](https://golang.org/dl/)
- **Google Cloud Text-to-Speech API key** - [Get one here](https://cloud.google.com/text-to-speech)
- **FFmpeg** - [Installation guide](https://ffmpeg.org/download.html)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/esl-toolkit.git
   cd esl-toolkit
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up your API key:**
   Create a `.env` file in the project root:
   ```env
   GOOGLE_API_KEY=your_google_tts_api_key_here
   ```

## Usage

### Spelling Practice Tool

The spelling practice tool randomly selects words from a word list and plays them using TTS, helping you practice spelling by listening.

**Setup:**
1. Create a `word_list.txt` file with one word per line:
   ```
   apple
   beautiful
   computer
   dictionary
   elephant
   ```

2. Open `main.go` and uncomment the spelling practice line:
   ```go
   func main() {
       cmd.RunSpellingPractice()
       // cmd.RunArticleTTS()
   }
   ```

3. Run the tool:
   ```bash
   go run main.go
   ```

**What happens:**
- Plays each word twice with a 3-second pause between repetitions
- Shows the word on screen after playing it
- Processes 5 words by default (configurable in `config/config.go`)

### Article TTS Tool

The article TTS tool reads entire text articles aloud, perfect for listening practice and pronunciation improvement.

**Setup:**
1. Create an `article.txt` file with your text:
   ```
   The quick brown fox jumps over the lazy dog.
   This is a sample article for TTS practice.
   ```

2. Open `main.go` and uncomment the article TTS line:
   ```go
   func main() {
       // cmd.RunSpellingPractice()
       cmd.RunArticleTTS()
   }
   ```

3. Run the tool:
   ```bash
   go run main.go
   ```

**What happens:**
- Converts the entire article to speech
- Saves the audio as `article.mp3`
- Plays the audio automatically
- Uses a random voice for variety

### MP3 Processing Script

The MP3 processing script helps you adjust the speed of existing audio files for different learning paces.

**Setup:**
1. Edit `scripts/process_mp3.sh`:
   ```bash
   # Set your input directories
   INPUT_DIRS=("/path/to/your/audio/files")

   # Set your output directory
   OUTPUT_DIR="/path/to/output/directory"

   # Adjust speed factor (1.15 = 15% faster)
   SPEED_FACTOR="1.15"
   ```

2. Run the script:
   ```bash
   cd scripts
   bash process_mp3.sh
   ```

## Configuration

All settings are in `config/config.go`:

```go
const (
    ArticleTTSFile                = "article.txt"           // Input file for articles
    ArticleTTSSpeechRate          = 1.0                     // Speech rate (1.0 = normal)
    SpellingPracticeFile          = "word_list.txt"         // Input file for words
    SpellingPracticeWords         = 5                       // Number of words to practice
    SpellingPracticePauseDuration = 3                       // Pause between repetitions (seconds)
)
```

## Project Structure

```
esl-toolkit/
├── cmd/                    # Main command implementations
│   ├── spelling.go         # Spelling practice tool
│   └── article_tts.go      # Article TTS tool
├── config/                 # Configuration management
│   └── config.go          # Settings and constants
├── internal/               # Internal packages
│   ├── audio/             # Audio playback functionality
│   ├── tts/               # Text-to-Speech integration
│   └── words/             # Word list management
├── scripts/                # Utility scripts
│   └── process_mp3.sh     # MP3 processing script
├── main.go                 # Main entry point
└── README.md              # This file
```

## Customization

### Adding New Voices

The TTS system uses Google's Text-to-Speech API. You can modify voice selection in `internal/tts/google.go`.

### Adjusting Speech Rate

For articles, modify `ArticleTTSSpeechRate` in `config/config.go`:
- `0.5` = 50% slower
- `1.0` = normal speed
- `1.5` = 50% faster

### Changing Practice Settings

Modify these constants in `config/config.go`:
- `SpellingPracticeWords`: Number of words to practice
- `SpellingPracticePauseDuration`: Pause between word repetitions



## License
[MIT License](LICENSE)


---

**Happy learning! 🎓**