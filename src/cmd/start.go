package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:           "start <service-name>",
	Short:         helpStartCmd,
	Long:          out.Banner(helpStartCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("start command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
