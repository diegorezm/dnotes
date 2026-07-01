package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for dnotes.
var rootCmd = &cobra.Command{
	Use:   "dnotes",
	Short: "A scriptable note-taking CLI for developer workflows",
	Long: `dnotes is a scriptable command-line note-taking tool designed to
integrate cleanly with editors and other developer tools (like Neovim).

It focuses on being easy to call from scripts, plugins, and keybindings,
so you can create, search, and manage notes without leaving your terminal
or editor. Every command is built to be composable with pipes, exit codes,
and structured output, making dnotes a good building block for custom
note-taking workflows rather than a closed, monolithic app.`,
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen
// once to rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags apply to rootCmd and every subcommand.
	// Good candidates: --config, --verbose, --json (for scriptable output).
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnotes.yaml)")
}
