package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

//go:embed yo.sh
var scripts embed.FS

func main() {
	var rootCmd = &cobra.Command{
		Use:   "command-runner",
		Short: "Command Runner is a simple tool to run a script with subcommands",
	}

	var cmdHello = &cobra.Command{
		Use:   "hello",
		Short: "Run the script with 'hello' argument",
		Run: func(cmd *cobra.Command, args []string) {
			runScript("hello", args)
		},
	}

	var cmdDawg = &cobra.Command{
		Use:   "dawg",
		Short: "Run the script with 'dawg' argument",
		Run: func(cmd *cobra.Command, args []string) {
			runScript("dawg", args)
		},
	}

	rootCmd.AddCommand(cmdHello, cmdDawg)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runScript(command string, args []string) {
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
