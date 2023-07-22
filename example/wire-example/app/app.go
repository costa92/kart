package app

import (
	"github.com/google/wire"
	"kart-io/kart/example/wire-example/config"
	"kart-io/kart/example/wire-example/options"
	"kart-io/kart/internal/command"
)

const commandDesc = `The Kart API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more api server information at:
    https://github.com/costa92/kart`

var ProviderHttpSeverSet = wire.NewSet(NewConfig)

func NewConfig() *command.App {
	opts := options.NewOptions()
	a := command.NewApp(
		"kart",
		// command.WithNoConfig(),
		command.WithOptions(opts),
		command.WithDescription(commandDesc),
		command.WithRunFunc(run(opts)),
	)
	return a
}

func run(opts *options.Options) command.RunFunc {
	return func(basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
