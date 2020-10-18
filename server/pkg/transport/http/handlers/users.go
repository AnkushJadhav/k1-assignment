package handlers

import (
	"net/http"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/modules/users"
	"github.com/AnkushJadhav/k1-assignment/server/utils"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/gin-gonic/gin"
)

// CreateUserHandler provides the HTTP handler for creating a user
func CreateUserHandler(store persistance.Client) gin.HandlerFunc {
	type body struct {
		Name     string `json:"name" validate:"required"`
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

		_, err = users.Create(store, inp.Name, inp.Email, inp.Password)
		if err != nil {
			if err, ok := err.(*users.ServiceError); ok {
				c.AbortWithStatusJSON(http.StatusOK, utils.JSONResponse(false, err.Error(), nil))
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}
