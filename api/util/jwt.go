package util

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var secret = os.Getenv("SECRET_KEY")

const (
	userIdKey         = "user_id"
	organizationIdKey = "organization_id"
	iatKey            = "iat"
	lifetime          = 30 * time.Minute
)

type Auth struct {
	UserID string
	Iat    int64
	Exp    int64
}

func GenerateToken(userId uuid.UUID) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIdKey: userId,
		iatKey:    now.Unix(),
	})

	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func ParseToken(signedString string) (*Auth, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired", signedString, err)
			} else {
				return nil, fmt.Errorf("%s is invalid", signedString, err)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid", signedString, err)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s:", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}

	userId := claims[userIdKey]
	iat := claims[iatKey]

	// reflectedValue := reflect.ValueOf(userId)
	// id, _ := reflectedValue.Interface().(uuid.UUID)

	// fmt.Println(userId)
	// fmt.Println(userId.(string))

	return &Auth{
		UserID: userId.(string),
		Iat:    int64(iat.(float64)),
	}, nil
}
