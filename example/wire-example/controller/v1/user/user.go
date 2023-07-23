package user

import "github.com/google/wire"

var ProvideUserAPISet = wire.NewSet(NewUserController)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}
