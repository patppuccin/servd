package cmd

import (
	"os"
	"servd/src/out"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:           "stop <service-name>",
	Short:         helpStopCmd,
	Long:          out.Banner(helpStopCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRun: func(cmd *cobra.Command, args []string) {
		if flagGlobalNoInteractive && len(args) == 0 {
			out.Error("No service name provided")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		out.Info("stop command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
