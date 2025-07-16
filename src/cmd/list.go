package cmd

import (
	"servd/src/out"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <service-name-pattern|optional>",
	Short: helpListCmd,
	Long:  out.Banner(helpListCmd),
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("list command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
