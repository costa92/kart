// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"kart-io/kart/example/wire-example/app"
)

// Injectors from wire.go:

func wireApp() (*app.ApiServer, func(), error) {
	options := app.NewOptionConfig()
	apiServer := app.NewApiServer(options)
	return apiServer, func() {
	}, nil
}
