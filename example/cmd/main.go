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
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

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
		fmt.Println("config", cfg)
		return nil
		//return Run(cfg)
	}
}
