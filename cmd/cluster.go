package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	internalsCluster "github.com/diegorezm/dnotes/internals/cluster"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage dnotes clusters",
	Long:  "Initialize and manage dnotes clusters (workspaces).",
}

var clusterInitCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize a dnotes cluster",
	Long: `Initialize a new dnotes cluster in the specified directory.

If no directory is provided, the current working directory is used.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var dir string

		if len(args) == 1 {
			dir = args[0]
		} else {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}
			dir = cwd
		}

		abs, err := filepath.Abs(dir)
		if err != nil {
			return fmt.Errorf("failed to resolve directory %s: %w", dir, err)
		}

		if err := internalsCluster.Init(abs); err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Initialized cluster in:", abs)
		return nil
	},
}

var clusterUseCmd = &cobra.Command{
	Use:   "use [directory]",
	Short: "Set the active dnotes cluster",
	Long: `Set the active dnotes cluster in the global config.

The directory must contain a .dnotes/ folder to be considered a valid cluster.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		abs, err := filepath.Abs(dir)
		if err != nil {
			return fmt.Errorf("failed to resolve path %s: %w", dir, err)
		}

		if err := internalsCluster.Use(abs); err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Active cluster set to:", abs)
		return nil
	},
}

func init() {
	clusterCmd.AddCommand(clusterUseCmd)
	clusterCmd.AddCommand(clusterInitCmd)
	rootCmd.AddCommand(clusterCmd)
}
