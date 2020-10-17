package users

import (
	"fmt"
	"regexp"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/utils"
)

const (
	emailTpl = "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
)

// Create validates and adds a user to the system
func Create(db persistance.Client, name, email, password string) (*models.User, error) {
	var err error
	err = validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateEmail(email)
	if err != nil {
		return nil, err
	}
	err = validatePassword(password)
	if err != nil {
		return nil, err
	}

	// generate user id
	id, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Hits:     0,
		IsActive: true,
	}

	return db.CreateUser(user)
}

// Update updates the users details in the system
func Update(db persistance.Client, user *models.User) error {
	return db.UpdateUser(user)
}

// Delete removes a user from the system. Soft delete is performed
func Delete(db persistance.Client, user *models.User) error {
	user.IsActive = false
	return db.UpdateUser(user)
}

// GetDetails gets the user details by id
func GetDetails(db persistance.Client, id string) (*models.User, error) {
	return db.GetUserByID(id)
}

// GetMultiple gets multiple users based on criteria
func GetMultiple(db persistance.Client, pageIndex, pageSize int, sortTerms, searchTerms map[string]string) ([]models.User, error) {
	user := models.User{
		Name:     searchTerms["name"],
		Email:    searchTerms["email"],
		IsActive: true,
	}

	sort := make([]persistance.Sorter, 0)
	for _, term := range sortTerms {
		var sortStyle persistance.SortStyle
		if sortTerms[term] == "asc" {
			sortStyle = persistance.ASC
		} else {
			sortStyle = persistance.DESC
		}
		sort = append(sort, persistance.Sorter{
			Attr:      term,
			Direction: sortStyle,
		})
	}

	pg := persistance.Paginator{
		PageSize: pageSize,
		Index:    pageIndex,
	}

	return db.GetUsers(user, sort, pg)
}

// validateName checks whether name is valid
func validateName(name string) error {
	if utils.IsZeroOfUnderlyingType(name) {
		return fmt.Errorf("Name cannot be blank")
	}
	return nil
}

// validateEmail checks whether email id is valid
func validateEmail(email string) error {
	if utils.IsZeroOfUnderlyingType(email) {
		return fmt.Errorf("Email cannot be blank")
	}
	valid, _ := regexp.MatchString(emailTpl, email)
	if !valid {
		return fmt.Errorf("Invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if utils.IsZeroOfUnderlyingType(password) {
		return fmt.Errorf("Email cannot be blank")
	}
	if len(password) < 10 {
		return fmt.Errorf("Password has to be atleast 10 characters long")
	}
	return nil
}
