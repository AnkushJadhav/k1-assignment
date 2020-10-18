package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	validityHours = 24 * time.Hour
)

// Issuer is used to issue and validate JWT tokens
type Issuer struct {
	secret string
}

// NewIssuer creates a new issuer for issuing JWTs signed with secret
func NewIssuer(secret string) *Issuer {
	return &Issuer{
		secret: secret,
	}
}

// Issue generates a signed JWT token for id
func (i *Issuer) Issue(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": getExpiry(time.Now()).Unix(),
		"nbf": time.Now().Unix(),
	})

	return token.SignedString([]byte(i.secret))
}

// IsValid checks if token is valid JWT issued by issuer i
func (i *Issuer) IsValid(token string) (bool, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(i.secret), nil
	})
	if err != nil {
		return false, err
	}
	if t.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return false, err
	}
}

// getExpiry adds number of validity hours to t
func getExpiry(t time.Time) time.Time {
	return t.Add(validityHours)
}
