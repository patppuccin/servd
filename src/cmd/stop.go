package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:           "stop <service-name>",
	Short:         helpStopCmd,
	Long:          out.Banner(helpStopCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("stop command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
