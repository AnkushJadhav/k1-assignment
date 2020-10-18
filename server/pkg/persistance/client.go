package persistance

import (
	"github.com/AnkushJadhav/k1-assignment/server/pkg/models"
)

// SortStyle defines the sorting direction when querying multiple records
type SortStyle string

// SortStyles
const (
	ASC  SortStyle = "asc"
	DESC SortStyle = "desc"
)

// Sorter defines the attribute to sort by and it's style
type Sorter struct {
	Attr      string
	Direction SortStyle
}

// Paginator represents the pagination when fetching multiple records
type Paginator struct {
	PageSize int
	Index    int
}

// Client is an interface that the application data persistance provider must implement
type Client interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUsers(users []models.User) error
	GetUser(*models.User) (*models.User, error)
	GetUsers(filter models.User, sort []Sorter, pg Paginator) ([]models.User, error)
}
