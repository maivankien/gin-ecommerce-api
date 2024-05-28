package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> database
func (uc *UserController) GetUserByID(c *gin.Context) {
	uid := c.Param("uid")
	c.JSON(http.StatusOK, gin.H{
		"uid":     uid,
		"message": uc.userService.GetInfoUser(),
	})
}
