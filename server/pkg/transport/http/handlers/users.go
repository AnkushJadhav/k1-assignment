package handlers

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/gin-gonic/gin"
)

// CreateUser provides the HTTP handler for creating a user
func CreateUser(store persistance.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
