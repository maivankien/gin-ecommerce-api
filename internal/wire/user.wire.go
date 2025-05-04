//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/maivankien/go-ecommerce-api/internal/controller"
	"github.com/maivankien/go-ecommerce-api/internal/repo"
	"github.com/maivankien/go-ecommerce-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
