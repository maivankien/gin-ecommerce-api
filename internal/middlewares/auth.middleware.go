package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/maivankien/go-ecommerce-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
