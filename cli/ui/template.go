package ui

import (
	"strings"
	"unicode"

	"github.com/marcelohmdias/dotforge/cli/print"
	"github.com/spf13/cobra"
)

func HelpFunc(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("help") {
		cmd.Root().PersistentPreRun(cmd, args)
	}
	w := cmd.OutOrStdout()
	usage := cmd.Long
	if usage == "" {
		usage = cmd.Short
	}
	usage = TrimSpaces(usage)
	if usage != "" {
		print.Fprintln(w, usage)
		print.Fprintln(w, "")
	}
	if cmd.Runnable() || cmd.HasSubCommands() {
		print.Fprintln(w, cmd.UsageString())
	}
}

func TrimSpaces(out string) string {
	return strings.TrimRightFunc(out, unicode.IsSpace)
}
