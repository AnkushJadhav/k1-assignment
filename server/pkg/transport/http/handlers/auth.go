package handlers

import (
	"net/http"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/modules/auth"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/transport/http/jwt"
	"github.com/AnkushJadhav/k1-assignment/server/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthenticator creates a auth middleware to verify JWT
func JWTAuthenticator(iss *jwt.Issuer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authentication")
		if err != nil {
			c.AbortWithStatusJSON(401, utils.JSONResponse(false, nil, nil))
			return
		}
		valid, err := iss.IsValid(token)
		if err != nil {
			c.AbortWithError(401, err)
			return
		}
		if !valid {
			c.AbortWithStatusJSON(401, utils.JSONResponse(false, "invalid token", nil))
			return
		}
		c.Next()
	}
}

// LoginHandler provides the HTTP handler for login
func LoginHandler(store persistance.Client, iss *jwt.Issuer) gin.HandlerFunc {
	type body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=10"`
	}
	return func(c *gin.Context) {
		inp := &body{}
		err := c.BindJSON(inp)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if ok, errs := utils.ValidateData(*inp); !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.JSONResponse(false, nil, errs))
			return
		}

		user, err := auth.GetUserByEmail(store, inp.Email)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if utils.IsZeroOfUnderlyingType(user) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.JSONResponse(false, "invalid email or password", nil))
			return
		}

		token, err := iss.Issue(user.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.SetCookie("Authentication", token, 0, "/", "", false, true)
		c.Status(http.StatusOK)
		return
	}
}

// LogoutHandler logs the user out by deleting the auth cookie
func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Authentication")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.SetCookie("Authentication", cookie, -1, "/", "", false, true)
		c.Status(http.StatusOK)
		return
	}
}
