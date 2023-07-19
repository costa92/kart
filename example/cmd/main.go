package main

import (
	"fmt"
	"kart-io/kart/example/cmd/config"
	"kart-io/kart/example/cmd/pkg"
)

const commandDesc = `The Kart API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/costa92/kart`

func main() {
	opts := pkg.NewOptions()
	application := pkg.NewApp(
		"kart",
		pkg.WithOptions(opts),
		pkg.WithDescription(commandDesc),
		pkg.WithRunFunc(run(opts)),
	)
	application.Run()
}

func run(opts *pkg.Options) pkg.RunFunc {
	return func(basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}

func Run(config *config.Config) error {
	fmt.Println(config)
	return nil
}
