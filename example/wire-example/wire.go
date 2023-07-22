//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	app2 "kart-io/kart/example/wire-example/app"
	"kart-io/kart/internal/command"
)

func wireApp() (*command.App, func(), error) {
	panic(wire.Build(app2.ProviderHttpSeverSet))
}
