package scriptrunner

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed yo.sh
var scripts embed.FS

func RunScript(command string, args []string) {
	scriptContent, err := scripts.ReadFile("yo.sh")
	if err != nil {
		fmt.Println("Error reading script:", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "yo-*.sh")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(scriptContent); err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}
	if err := tmpFile.Close(); err != nil {
		fmt.Println("Error closing temp file:", err)
		return
	}

	cmdArgs := append([]string{command}, args...)
	cmd := exec.Command("/bin/bash", append([]string{tmpFile.Name()}, cmdArgs...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running script:", err)
	}
}
