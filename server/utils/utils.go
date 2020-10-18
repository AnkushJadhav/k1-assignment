package utils

import (
	"reflect"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
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

// ValidateData validates a struct using the validate tags on it's fields
func ValidateData(dataSet interface{}) (bool, map[string]string) {
	validate := validator.New()

	err := validate.Struct(dataSet)
	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		errors := make(map[string]string)
		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "The " + name + " is required"
				break
			case "email":
				errors[name] = "The " + name + " should be a valid email"
				break
			case "eqField":
				errors[name] = "The " + name + " should be equal to the " + err.Param()
				break
			default:
				errors[name] = "The " + name + " is invalid"
				break
			}
		}

		return false, errors
	}

	return true, nil
}

// JSONResponse wraps the response status and data with some rules
func JSONResponse(success bool, msg, data interface{}) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["success"] = success
	if msg != nil {
		resp["message"] = msg
	}
	if data != nil {
		resp["data"] = data
	}

	return resp
}
