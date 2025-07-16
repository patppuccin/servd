package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:           "register <service-name>",
	Short:         helpRegisterCmd,
	Long:          out.Banner(helpRegisterCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Aliases:       []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("register command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
