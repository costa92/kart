package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kart-io/kart/internal/command"
	"kart-io/kart/pkg/cliflag"
)

type App struct {
	basename    string
	options     command.CliOptions
	runFunc     RunFunc
	description string
	commands    []*command.Command
	cmd         *cobra.Command
	noConfig    bool
}

// Option defines optional parameters for initializing the application
// structure.
type Option func(*App)

// WithOptions returns a function that calls each Option in Options
func WithOptions(opt command.CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

// RunFunc defines the application's startup callback function.
type RunFunc func(basename string) error

// WithRunFunc returns a function that calls each RunFunc in RunFuncs
func WithRunFunc(runFunc RunFunc) Option {
	return func(a *App) {
		a.runFunc = runFunc
	}
}

// WithDescription is used to set the description of the application.
func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

func NewApp(basename string, opts ...Option) *App {
	a := &App{
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	a.buildCommand()
	return a
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           command.FormatBaseName(a.basename),
		Short:         "kart",
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// cmd.SetUsageTemplate(usageTemplate)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	command.InitFlags(cmd.Flags())

	if len(a.commands) > 0 {
		for _, c := range a.commands {
			cmd.AddCommand(c.CobraCommand())
		}
		cmd.SetHelpCommand(command.HelpCommand(command.FormatBaseName(a.basename)))
	}

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	command.PrintFlags(cmd.Flags())
	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}
	var namedFlagSets cliflag.NamedFlagSets
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}
	// run application
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}
	return nil
}

// Run is used to launch the application.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
