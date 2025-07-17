package cmd

import (
	"os"
	"servd/src/out"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:           "restart <service-name>",
	Short:         helpRestartCmd,
	Long:          out.Banner(helpRestartCmd),
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
		out.Info("restart command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
