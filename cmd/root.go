package cmd

import (
	"os"

	"github.com/marcelohmdias/dotforge/cli/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "A smart way to forge your system",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.Banner()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpFunc(ui.HelpFunc)
}
