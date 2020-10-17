package http

import (
	"github.com/gin-gonic/gin"
)

// NewJWTAuthenticator creates a auth middleware to verify JWT
func NewJWTAuthenticator(iss *Issuer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authentication")
		if err != nil {
			c.String(401, "Unauthorized request")
		}
		valid, err := iss.IsValid(token)
		if err != nil || !valid {
			c.String(401, "Unauthorized request")
		}
		c.Next()
	}
}
