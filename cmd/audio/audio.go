package audio

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Execute(from, to string) ([]byte, error) {
	return exec.Command(
		"ffmpeg",
		"-i", fmt.Sprintf("%s.mp4", from),
		"-vn",
		"-acodec", "mp3",
		fmt.Sprintf("%s.mp3", to),
	).CombinedOutput()
}

func FromVideo(videoPath, outputPath string) ([]byte, error) {
	return exec.Command(
		"ffmpeg",
		"-y",
		"-i", filepath.Clean(videoPath),
		"-vn",
		"-acodec", "libmp3lame",
		filepath.Clean(outputPath),
	).CombinedOutput()
}
