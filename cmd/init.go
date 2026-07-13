package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCMD = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize a dnotes cluster",
	Long: `Initialize a new dnotes cluster in the specified directory.

If no directory is provided, the current working directory is used.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var directory string

		if len(args) == 1 {
			directory = args[0]
		} else {
			var err error
			directory, err = os.Getwd()
			if err != nil {
				return err
			}
		}

		fmt.Println("Initializing cluster in:", directory)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCMD)
}
