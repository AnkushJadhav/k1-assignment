package users

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/utils"
)

const (
	emailTpl = "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
)

// ServiceError is the generic data type to hold business logic errors
type ServiceError struct {
	err string
}

func (e *ServiceError) Error() string {
	return e.err
}

// Create validates and adds a user to the system
func Create(db persistance.Client, user *models.User) (*models.User, error) {
	var err error
	err = validateName(user.Name)
	if err != nil {
		return nil, &ServiceError{err.Error()}
	}
	err = validateEmail(user.Email)
	if err != nil {
		return nil, &ServiceError{err.Error()}
	}
	err = validatePassword(user.Password)
	if err != nil {
		return nil, &ServiceError{err.Error()}
	}

	// generate user id
	id, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}

	exists, err := db.GetUser(&models.User{
		Email: user.Email,
	})
	if !utils.IsZeroOfUnderlyingType(exists) {
		return nil, &ServiceError{"user with this email already exists"}
	}

	user.ID = id
	user.Hits = 0

	return db.CreateUser(user)
}

// Update updates the users details in the system
func Update(db persistance.Client, user *models.User) error {
	return db.UpdateUser(user)
}

// DeleteMultiple removes users from the system
func DeleteMultiple(db persistance.Client, users []models.User) error {
	return db.DeleteUsers(users)
}

// GetDetails gets the user details by id
func GetDetails(db persistance.Client, id string) (*models.User, error) {
	return db.GetUser(&models.User{
		ID: id,
	})
}

// GetMultiple gets multiple users based on criteria
func GetMultiple(db persistance.Client, pageIndex string, pageSize int, sortTerms []string, searchTerms map[string]string) ([]models.User, error) {
	user := models.User{
		Name:  searchTerms["name"],
		Email: searchTerms["email"],
	}

	sort := make([]persistance.Sorter, 0)
	for _, term := range sortTerms {
		var sortStyle persistance.SortStyle
		split := strings.Split(term, ":")
		if len(split) == 1 {
			sortStyle = persistance.ASC
		} else if split[1] == "desc" {
			sortStyle = persistance.DESC
		} else {
			sortStyle = persistance.ASC
		}
		sort = append(sort, persistance.Sorter{
			Attr:      split[0],
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
