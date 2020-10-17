package utils

import (
	"reflect"

	"github.com/google/uuid"
)

// GenerateID generates a cryptographically random UUID
func GenerateID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// IsZeroOfUnderlyingType detects whether x is the zero value of it's type
func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
