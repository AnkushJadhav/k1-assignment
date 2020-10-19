package handlers

import (
	"net/http"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"

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

		u := &models.User{
			Name:     inp.Name,
			Email:    inp.Email,
			Password: inp.Password,
		}
		_, err = users.Create(store, u)
		if err != nil {
			if err, ok := err.(*users.ServiceError); ok {
				c.JSON(http.StatusOK, utils.JSONResponse(false, err.Error(), nil))
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

// GetUserByIDHandler provides the HTTP handler for fetching a user
func GetUserByIDHandler(store persistance.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("id")
		if utils.IsZeroOfUnderlyingType(uid) {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.JSONResponse(false, "id not found", nil))
			return
		}

		u, err := users.GetDetails(store, uid)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if utils.IsZeroOfUnderlyingType(u) {
			c.JSON(http.StatusOK, utils.JSONResponse(false, "no user found", nil))
			return
		}

		c.JSON(http.StatusOK, utils.JSONResponse(true, nil, u))
		return
	}
}

// GetUsersHandler provides the HTTP handler for fetching multiple users with pagination, search and sort
func GetUsersHandler(store persistance.Client) gin.HandlerFunc {
	type query struct {
		Pagesize  int      `form:"limit" binding:"required"`
		Pageindex string   `form:"marker" binding:"-"`
		Sort      []string `form:"sort" binding:"-"`
		Name      string   `form:"name" binding:"-"`
		Email     string   `form:"email" binding:"-"`
	}
	return func(c *gin.Context) {
		var q query
		err := c.BindQuery(&q)
		if err != nil {
			return
		}

		st := make(map[string]string)
		st["name"] = q.Name
		st["email"] = q.Email

		u, err := users.GetMultiple(store, q.Pageindex, q.Pagesize, q.Sort, st)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, utils.JSONResponse(true, nil, u))
	}
}

// EditUserHandler provides the HTTP handler for editing a user
func EditUserHandler(store persistance.Client) gin.HandlerFunc {
	type body struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
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

		uid := c.Param("id")
		if utils.IsZeroOfUnderlyingType(uid) {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.JSONResponse(false, "id not found", nil))
			return
		}

		u := &models.User{
			ID:    uid,
			Name:  inp.Name,
			Email: inp.Email,
		}

		err = users.Update(store, u)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
		return
	}
}

// DeleteUser deletes a user with the given id
func DeleteUser(store persistance.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("id")
		if utils.IsZeroOfUnderlyingType(uid) {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.JSONResponse(false, "id not found", nil))
			return
		}

		u, err := users.GetDetails(store, uid)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if utils.IsZeroOfUnderlyingType(u) {
			c.JSON(http.StatusOK, utils.JSONResponse(false, "no user found", nil))
			return
		}

		uarr := make([]models.User, 0)
		uarr = append(uarr, *u)

		err = users.DeleteMultiple(store, uarr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
		return
	}
}