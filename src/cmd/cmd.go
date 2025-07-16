package cmd

import (
	"fmt"
	"os"
	"servd/src/constants"
	"servd/src/out"

	"github.com/spf13/cobra"
)

const (
	helpRootCmd      = "Cross-platform service management tool"
	helpRegisterCmd  = "Register a new system service"
	helpStartCmd     = "Start a registered service"
	helpStopCmd      = "Stop a running service"
	helpRestartCmd   = "Restart a running or stopped service"
	helpUninstallCmd = "Uninstall & remove a registered service"
	helpEditCmd      = "Edit the configuration of an existing registered service"
	helpListCmd      = "List registered services"
	helpDoctorCmd    = "Run tool diagnostics"
)

var (
	flagGlobalNoInteractive   bool
	flagRegisterDisplayName   string
	flagRegisterDescription   string
	flagRegisterTags          []string
	flagRegisterExec          string
	flagRegisterDir           string
	flagRegisterEnv           []string
	flagRegisterEnvFiles      []string
	flagRegisterStartupType   string
	flagRegisterRestartPolicy string
	flagRegisterRestartConfig string
	flagRegisterStartNow      bool
	flagEditDisplayName       string
	flagEditDescription       string
	flagEditTags              []string
	flagEditExec              string
	flagEditDir               string
	flagEditEnv               []string
	flagEditEnvFiles          []string
	flagEditStartupType       string
	flagEditRestartPolicy     string
	flagEditRestartConfig     string
	flagEditRestartNow        bool
	flagUninstallConfirm      bool
	flagListStatus            string
	flagListFormat            string
	flagListMode              string
	flagListExportPath        string
	flagDoctorFormat          string
	flagDoctorMode            string
	flagDoctorExportPath      string
)

var rootCmd = &cobra.Command{
	Use:           constants.AppAbbrName,
	Short:         helpRootCmd,
	Long:          out.Banner(helpRootCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       constants.AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var validListOutputFormats = []string{"json", "yaml", "csv", "table"}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVar(&flagGlobalNoInteractive, "no-interactive", false, "disable interactive mode")

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Help()
		fmt.Println()
		out.Error("Error parsing the provided flags: " + err.Error())
		os.Exit(1)
		return nil
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		out.Error(err.Error())
		os.Exit(1)
	}
}
