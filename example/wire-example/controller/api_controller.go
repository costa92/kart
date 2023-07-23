package controller

import (
	"github.com/google/wire"
)

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(ProvideApiController)

type ApiController struct {
}

func ProvideApiController() *ApiController {
	return &ApiController{}
}
