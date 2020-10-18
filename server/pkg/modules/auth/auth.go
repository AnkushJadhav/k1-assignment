package auth

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
)

// GetUserByEmail checks whether a user with email exists in the system
func GetUserByEmail(db persistance.Client, email string) (*models.User, error) {
	return db.GetUser(&models.User{
		Email: email,
	})
}
