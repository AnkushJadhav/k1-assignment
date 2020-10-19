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

	db = db.Debug()

	return &Client{db}, nil
}

// CreateUser creates a new user in the users table
func (db *Client) CreateUser(user *models.User) (*models.User, error) {
	result := db.Create(user)
	return user, result.Error
}

// UpdateUser edits the user in the users table
func (db *Client) UpdateUser(user *models.User) error {
	result := db.Updates(user)
	return result.Error
}

// DeleteUsers deletes a user from the users table
func (db *Client) DeleteUsers(users []models.User) error {
	result := db.Delete(&models.User{}, users)
	return result.Error
}

// GetUser fetches a user from the users table
func (db *Client) GetUser(query *models.User) (*models.User, error) {
	user := &models.User{}
	result := db.Where(query).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

// GetUsers fetches multiple users from the user table with optional filter
func (db *Client) GetUsers(filter models.User, sort []persistance.Sorter, pg persistance.Paginator) ([]models.User, error) {
	tx := db.DB
	// prepare filter statments
	if !utils.IsZeroOfUnderlyingType(filter.Name) {
		tx = tx.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if !utils.IsZeroOfUnderlyingType(filter.Email) {
		tx = tx.Where("email LIKE ?", "%"+filter.Email+"%")
	}

	// prepare sort statements
	for _, s := range sort {
		tx = tx.Order(s.Attr + " " + string(s.Direction))
	}
	tx = tx.Order("id asc")

	// prepare pagination
	if pg.Index != "" {
		tx = tx.Where("id > ?", pg.Index)
	}
	tx = tx.Limit(pg.PageSize)

	var users []models.User
	result := tx.Find(&users)
	return users, result.Error
}
