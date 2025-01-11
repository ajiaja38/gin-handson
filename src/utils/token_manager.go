package utils

import (
	"errors"
	"res-gin/src/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, duration time.Duration, secret []byte) (string, error) {
	claims := &model.UserClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(token string, secret []byte, claims *model.UserClaims) error {
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})

	if err != nil || !parsedToken.Valid {
		return errors.New("Invalid Token")
	}

	return nil
}
