# ESL Toolkit

A collection of small, practical tools for English as a Second Language (ESL) learners. This toolkit is designed to be simple, modular, and developer-friendly.

> **Note**: This project is not currently designed as a CLI app. To use a tool, uncomment the corresponding function call in `main.go`.

## Features

- **Spelling Practice**: Plays random words using Text-to-Speech (TTS) for spelling practice.
- **Article TTS**: Reads an article aloud using TTS and saves/plays the audio.
- **MP3 Processing Script**: Batch-processes MP3 files to adjust playback speed using FFmpeg.


## Getting Started

### Prerequisites

- Go 1.18+ installed
- A Google Cloud Text-to-Speech API key
- [ffmpeg](https://ffmpeg.org/) installed (for audio playback and scripts)

### Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/esl-toolkit.git
   cd esl-toolkit
   ```

2. **Set up your environment variables:**
   - Create a `.env` file in the project root:
     ```
     GOOGLE_API_KEY=your_google_tts_api_key
     ```
   - Or export it in your shell:
     ```sh
     export GOOGLE_API_KEY=your_google_tts_api_key
     ```

3. **Install Go dependencies:**
   ```sh
   go mod tidy
   ```

### Usage

#### Spelling Practice or Article TTS

1. Open main.go and uncomment the desired function:
   - `cmd.RunSpellingPractice()`
   - `cmd.RunArticleTTS()`

2. Run the project:
    ```sh
    go run main.go
    ```

#### MP3 Processing Script

1. Edit `process_mp3.sh` to set your input/output directories.

2. Execute the script:
    ```sh
    cd scripts
    bash process_mp3.sh
    ```

## Configuration

- All configuration is managed via environment variables and the `config` package.
- ⚠️ **Never commit your API keys or secrets to version control.**
  The `.env` file is ignored by `.gitignore`.


## License

See [LICENSE](LICENSE) for details.