package audio

import (
	"os"
	"os/exec"
)

func PlayAudio(audioData []byte) error {
	tmpfile, err := os.CreateTemp("", "audio*.mp3")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(audioData); err != nil {
		return err
	}
	if err := tmpfile.Close(); err != nil {
		return err
	}

	cmd := exec.Command("afplay", tmpfile.Name())
	return cmd.Run()
}
