package cmd

import (
	"servd/src/models"
	"servd/src/out"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:           "register <service-name>",
	Short:         helpRegisterCmd,
	Long:          out.Banner(helpRegisterCmd),
	SilenceErrors: true,
	SilenceUsage:  true,
	Aliases:       []string{"add"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !flagGlobalNoInteractive {
			return
		}
		var errs []string
		if len(args) == 0 {
			errs = append(errs, "No service name provided")
		}
		if flagRegisterExec == "" {
			errs = append(errs, "No exec command provided (required)")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&flagRegisterDisplayName, "display-name", "n", "", "display name for the service")
	registerCmd.Flags().StringVarP(&flagRegisterDescription, "description", "d", "", "service description")
	registerCmd.Flags().StringSliceVarP(&flagRegisterTags, "tags", "t", nil, "list of tags (comma-separated)")

	registerCmd.Flags().StringVarP(&flagRegisterExec, "exec", "x", "", "command to run (with args)")
	registerCmd.Flags().StringVarP(&flagRegisterDir, "dir", "w", "", "working directory (defaults to curr dir)")
	registerCmd.Flags().StringSliceVarP(&flagRegisterEnv, "env", "e", nil, "list of env vars (e.g. KEY=VALUE)")
	registerCmd.Flags().StringSliceVarP(&flagRegisterEnvFiles, "env-files", "f", nil, "list of path(s) to .env files")

	registerCmd.Flags().StringVarP(&flagRegisterStartupType, "startup-type", "s", "auto", "startup type ("+strings.Join(validStartupTypes, ",")+")")
	registerCmd.Flags().StringVarP(&flagRegisterRestartPolicy, "restart-policy", "r", "", "restart policy ("+strings.Join(validRestartPolicies, ",")+")")
	registerCmd.Flags().StringVarP(&flagRegisterRestartConfig, "restart-config", "c", "", "restart config (max=3;delay=5s;mult=2.0)")

	registerCmd.Flags().BoolVar(&flagRegisterStartNow, "start", false, "start service after register")
}

func ShowRegisterForm() (*models.ServiceSpec, error) {
	var input models.ServiceSpec

	var tagsList []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Display Name").
				Placeholder("My Worker Service").
				Value(&input.DisplayName).
				Validate(huh.ValidateNotEmpty()),

			huh.NewInput().
				Title("Description").
				Placeholder("Describe what the service does").
				Value(&input.Description),

			huh.NewInput().
				Title("Tags").
				Description("Comma-separated tags").
				Placeholder("worker,batch,critical").
				Value(tagsList),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("Command to Execute").
				Placeholder("/usr/bin/my-app --run").
				Value(&input.Exec).
				Validate(huh.ValidateNotEmpty()),

			huh.NewInput().
				Title("Working Directory").
				Placeholder("/opt/my-app").
				Value(&input.Dir),

			huh.NewInput().
				Title("Environment Variables").
				Description("Comma-separated list like KEY=value").
				ValueFunc(func(s string) error {
					input.EnvVars = strings.Split(s, ",")
					return nil
				}),

			huh.NewInput().
				Title("Env Files").
				Description("Comma-separated paths to .env files").
				ValueFunc(func(s string) error {
					input.EnvFiles = strings.Split(s, ",")
					return nil
				}),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Startup Type").
				Options(
					huh.NewOption("Auto", "auto"),
					huh.NewOption("Manual", "manual"),
				).
				Value(&input.StartupType).
				Validate(huh.ValidateOneOf(validStartupTypes...)),

			huh.NewSelect[string]().
				Title("Restart Policy").
				Options(
					huh.NewOption("Always", "always"),
					huh.NewOption("On Failure", "on-fail"),
					huh.NewOption("Until Stopped", "till-stop"),
				).
				Value(&input.RestartPolicy).
				Validate(huh.ValidateOneOf(validRestartPolicies...)),

			huh.NewInput().
				Title("Restart Config").
				Description("e.g. max=3;delay=5s;mult=2.0").
				Value(&input.RestartConfig),

			huh.NewConfirm().
				Title("Start Immediately After Register?").
				Value(&input.StartNow),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return &input, nil
}
