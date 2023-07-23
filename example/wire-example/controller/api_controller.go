package controller

import "kart-io/kart/example/wire-example/controller/v1/user"

type ApiController struct {
	UserController *user.UserController
}

func ProvideApiController() *ApiController {
	return &ApiController{
		UserController: user.NewUserController(),
	}
}
