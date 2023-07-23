//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"kart-io/kart/example/wire-example/app"
)

func wireApp() (*app.ApiServer, func(), error) {
	panic(wire.Build(
		app.ProviderApiServerSet,
	))
}
