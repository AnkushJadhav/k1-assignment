package http

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/transport/http/handlers"
	"github.com/gin-gonic/gin"
)

func (s *Server) setAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", handlers.LoginHandler(s.store, s.issuer))
	router.POST("/logout", handlers.JWTAuthenticator(s.issuer))
}

func (s *Server) setAPIRoutes(router *gin.RouterGroup) {
	router.Use(handlers.JWTAuthenticator(s.issuer))

	router.POST("/user", handlers.CreateUserHandler(s.store))
}
