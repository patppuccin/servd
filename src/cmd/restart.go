package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:           "restart <service-name>",
	Short:         helpRestartCmd,
	Long:          out.Banner(helpRestartCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("restart command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
