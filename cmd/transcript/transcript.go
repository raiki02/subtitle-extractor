package transcript

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Execute(audioName string) ([]byte, error) {
	return exec.Command(
		"whisper-cli",
		"-m", "./models/ggml-small.bin",
		"-f", fmt.Sprintf("%s.mp3", audioName),
		"--prompt", "以下是普通话的句子，请使用简体中文输出。",
		//"-otxt",
		"-oj",
		"-l", "zh",
	).CombinedOutput()
}

func Text(audioPath, modelPath string) (string, []byte, error) {
	outputBase := filepath.Clean(audioPath)
	outputPath := outputBase + ".txt"
	out, err := exec.Command(
		"whisper-cli",
		"-m", filepath.Clean(modelPath),
		"-f", filepath.Clean(audioPath),
		"--prompt", "以下是普通话的句子，请使用简体中文输出。",
		"-otxt",
		"-l", "zh",
	).CombinedOutput()
	return outputPath, out, err
}
