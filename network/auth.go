package network

import (
	"github.com/gin-gonic/gin"
	"go_grpc/types"
	"net/http"
)

func (n *Network) login(c *gin.Context) {
	var req types.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := n.gRPCClient.CreateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": res})
}

func (n *Network) verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
