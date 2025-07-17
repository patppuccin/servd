package cmd

import (
	"servd/src/out"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:           "edit <service-name>",
	Short:         helpEditCmd,
	Long:          out.Banner(helpEditCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		out.Info("edit command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&flagEditDisplayName, "display-name", "n", "", "display name for the service")
	editCmd.Flags().StringVarP(&flagEditDescription, "description", "d", "", "service description")
	editCmd.Flags().StringSliceVarP(&flagEditTags, "tags", "t", nil, "list of tags (comma-separated)")

	editCmd.Flags().StringVarP(&flagEditExec, "exec", "x", "", "command to run (with args)")
	editCmd.Flags().StringVarP(&flagEditDir, "dir", "w", "", "working directory")
	editCmd.Flags().StringSliceVarP(&flagEditEnv, "env", "e", nil, "list of env vars (e.g. KEY=VALUE)")
	editCmd.Flags().StringSliceVarP(&flagEditEnvFiles, "env-files", "f", nil, "list of path(s) to .env files")

	editCmd.Flags().StringVarP(&flagEditStartupType, "startup-type", "s", "", "startup type ("+strings.Join(validStartupTypes, ",")+")")
	editCmd.Flags().StringVarP(&flagEditRestartPolicy, "restart-policy", "r", "", "restart policy ("+strings.Join(validRestartPolicies, ",")+")")
	editCmd.Flags().StringVarP(&flagEditRestartConfig, "restart-config", "c", "", "restart config (max=3;delay=5s;mult=2.0)")

	editCmd.Flags().BoolVar(&flagEditRestartNow, "restart", false, "restart service after editing")

}
