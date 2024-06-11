package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/internal/service"
	"github.com/maivankien/go-ecommerce-api/pkg/response"
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
	response.SuccessResponse(c, http.StatusOK, uc.userService.GetInfoUser())
}
