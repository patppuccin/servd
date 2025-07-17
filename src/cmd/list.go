package cmd

import (
	"os"
	"servd/src/out"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:           "list <service-name-pattern>",
	Short:         helpListCmd,
	Long:          out.Banner(helpListCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRun: func(cmd *cobra.Command, args []string) {
		var errs []string
		if !slices.Contains(validListStatuses, flagListStatus) {
			errs = append(errs, "Invalid status filter: "+flagListStatus)
		}
		if !slices.Contains(validListOutputFormats, flagListFormat) {
			errs = append(errs, "Invalid output format: "+flagListFormat)
		}
		if !slices.Contains(validListModes, flagListMode) {
			errs = append(errs, "Invalid run mode: "+flagListMode)
		}
		if len(errs) > 0 {
			for _, err := range errs {
				out.Error(err)
			}
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		out.Info("list command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&flagListStatus, "status", "s", "all", "status filter ("+strings.Join(validListStatuses, ",")+")")
	listCmd.Flags().StringVarP(&flagListFormat, "format", "f", "table", "out format ("+strings.Join(validListOutputFormats, ",")+")")
	listCmd.Flags().StringVarP(&flagListMode, "mode", "m", "simple", "run mode ("+strings.Join(validListModes, ",")+")")
	listCmd.Flags().StringVarP(&flagListExportPath, "export", "o", "", "path to the (optional) output file")
}
