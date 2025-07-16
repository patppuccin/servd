package cmd

import (
	"os"
	"servd/src/out"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var validDoctorModes = []string{"quick", "full", "detailed"}
var validDoctorOutputFormats = []string{"json", "yaml", "csv", "list"}
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: helpDoctorCmd,
	Long:  out.Banner(helpDoctorCmd),
	PreRun: func(cmd *cobra.Command, args []string) {
		var errs []string
		if !slices.Contains(validDoctorOutputFormats, flagDoctorFormat) {
			errs = append(errs, "Invalid output format: "+flagDoctorFormat)
		}
		if !slices.Contains(validDoctorModes, flagDoctorMode) {
			errs = append(errs, "Invalid run mode: "+flagDoctorMode)
		}
		for _, err := range errs {
			out.Error(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		out.Info("doctor command is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
	doctorCmd.Flags().StringVarP(&flagDoctorFormat, "format", "f", "list", "output format (options: "+strings.Join(validDoctorOutputFormats, ", ")+")")
	doctorCmd.Flags().StringVarP(&flagDoctorMode, "mode", "m", "quick", "run mode (options: "+strings.Join(validDoctorModes, ", ")+")")
	doctorCmd.Flags().StringVarP(&flagDoctorExportPath, "export-path", "o", "", "path to the (optional) output file")
}
