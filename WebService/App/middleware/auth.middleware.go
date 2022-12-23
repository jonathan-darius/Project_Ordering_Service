package middleware

import (
	gRPCFunc "WebService/App/gRPC_Configs/User"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := gRPCFunc.ValidateToken(clientToken)

		if !claims.Status {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("uid", claims.UserId)
		c.Set("Messages", claims.Messages)
		c.Set("role", claims.Role)
		c.Next()
	}
}
