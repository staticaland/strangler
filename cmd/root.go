package cmd

import (
	"github.com/spf13/cobra"
	"github.com/staticaland/strangler/scriptrunner"
)

func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "command-runner",
		Short: "Command Runner is a simple tool to run a script with subcommands",
		// ... other setup if needed ...
	}

	// setup subcommands
	rootCmd.AddCommand(newHelloCommand(), newDawgCommand())

	return rootCmd
}

func newHelloCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Run the script with 'hello' argument",
		Run: func(cmd *cobra.Command, args []string) {
			scriptrunner.RunScript("hello", args)
		},
	}
}

func newDawgCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "dawg",
		Short: "Run the script with 'dawg' argument",
		Run: func(cmd *cobra.Command, args []string) {
			scriptrunner.RunScript("dawg", args)
		},
	}
}
