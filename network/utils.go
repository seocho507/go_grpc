package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 1. Get Bearer token
		token := getAuthToken(context)

		if strings.EqualFold(token, "") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			context.Abort()
			return
		}

		// 2. Validate token
		_, err := n.gRPCClient.ValidateToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			context.Abort()
			return
		}

		context.Next()
	}
}

func getAuthToken(context *gin.Context) string {

	var token string

	// Bearer ~~~
	header := context.GetHeader("Authorization")

	authSliced := strings.Split(header, " ")
	if len(authSliced) > 1 {
		token = authSliced[1]
		return token
	}

	return token
}
