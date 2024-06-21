package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/internal/controller"
	"github.com/maivankien/go-ecommerce-api/internal/middlewares"
)

func Authen() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before Authen middleware")
		c.Next()
		fmt.Println("Alter Authen middleware 2")
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before Logger middleware")
		c.Next()
		fmt.Println("Alter Logger middleware 2")
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenMiddleware(), Logger())

	v1 := r.Group("/v1")
	{
		v1.GET("/user/:uid", controller.NewUserController().GetUserByID)
	}

	return r
}
