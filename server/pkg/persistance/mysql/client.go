package mysql

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client represents a client to a mysql database
type Client struct {
	*gorm.DB
}

// New creates a new client after initiating a connection to the database located at dsn
func New(dsn string) (*Client, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}

// CreateUser creates a new user in the users table
func (db *Client) CreateUser(user *models.User) (*models.User, error) {
	result := db.Create(user)
	return user, result.Error
}

// UpdateUser edits the user in the users table
func (db *Client) UpdateUser(user *models.User) error {
	result := db.Save(user)
	return result.Error
}

// DeleteUser deletes a user from the users table
func (db *Client) DeleteUser(user *models.User) error {
	result := db.Delete(user)
	return result.Error
}

// GetUser fetches a user from the users table
func (db *Client) GetUser(user *models.User) (*models.User, error) {
	result := db.First(user)
	return user, result.Error
}

// GetUsers fetches multiple users from the user table with optional filter
func (db *Client) GetUsers(filter models.User, sort []persistance.Sorter, pg persistance.Paginator) ([]models.User, error) {
	// prepare filter statments
	if !utils.IsZeroOfUnderlyingType(filter.Name) {
		db.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if !utils.IsZeroOfUnderlyingType(filter.Email) {
		db.Where("email LIKE ?", "%"+filter.Email+"%")
	}
	db.Where("is_active = ?", filter.IsActive)

	// prepare sort statements
	for _, s := range sort {
		db.Order(s.Attr + " " + string(s.Direction))
	}

	// prepare pagination
	offset := pg.PageSize * (pg.Index + 1)
	db.Offset(offset)
	db.Limit(pg.PageSize)

	var users []models.User
	result := db.Find(users)
	return users, result.Error
}