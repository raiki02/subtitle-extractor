package download

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Execute(url string, name string) ([]byte, error) {
	return exec.Command("yt-dlp",
		"-o", fmt.Sprintf("%s.%%(ext)s", name),
		url).CombinedOutput()
}

func Video(url, outputPath string) ([]byte, error) {
	return exec.Command(
		"yt-dlp",
		"-f", "bv*+ba/b",
		"--merge-output-format", "mp4",
		"-o", filepath.Clean(outputPath),
		url,
	).CombinedOutput()
}
