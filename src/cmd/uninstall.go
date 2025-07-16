package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("uninstall command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
