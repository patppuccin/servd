package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:           "edit <service-name>",
	Short:         helpEditCmd,
	Long:          out.Banner(helpEditCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("edit command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
