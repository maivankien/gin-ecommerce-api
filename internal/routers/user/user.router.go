package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	// This is non dependency
	// userRepo := repo.NewUserRepository()
	// userService := service.NewUserService(userRepo)
	// userController := controller.NewUserController(userService)

	// Wire
	// Dependency Injection (DI)
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/opt")
	}

	// Private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/profile")
	}
}
