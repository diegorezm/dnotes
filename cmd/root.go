package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for dnotes.
var rootCmd = &cobra.Command{
	Use:   "dnotes",
	Short: "A scriptable note-taking CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

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
