package http

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/gin-gonic/gin"
)

func setInternalRoutes(router *gin.RouterGroup, store persistance.Client) {
	router.POST("/user")
}
