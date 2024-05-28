package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/user/:uid", controller.NewUserController().GetUserByID)
	}

	return r
}
