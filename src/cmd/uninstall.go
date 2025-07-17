package cmd

import (
	"os"
	"servd/src/out"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:           "uninstall <service-name>",
	Short:         helpUninstallCmd,
	Long:          out.Banner(helpUninstallCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Aliases:       []string{"remove", "rm", "delete", "del"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if flagGlobalNoInteractive && len(args) == 0 {
			out.Error("No service name provided")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		out.Info("uninstall command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
