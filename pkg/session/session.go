package session

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var tokenKey = []byte("fvoNImvpdms023sv0s9vs")

type UserClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type Claims struct {
	User UserClaims `json:"user"`
	jwt.StandardClaims
}

type JWTSessionsManager struct{}

func (jsm JWTSessionsManager) GetUser(inToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(inToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, errors.Errorf("bad sign method session token, expected: \"HS256\", got: \"%s\"", method.Alg())
		}
		return tokenKey, nil
	})

	if err != nil || !token.Valid {
		return -1, "", errors.Wrapf(err, "can`t parse session token \"%s\"", inToken)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return -1, "", errors.Errorf("can`t parse session token \"%s\"", inToken)
	}

	err = claims.Valid()
	if err != nil {
		return -1, "", errors.Wrap(err, "session isn`t valid")
	}

	return claims.User.ID, claims.User.Role, nil
}

func (jsm JWTSessionsManager) CreateSession(id int, role string) (string, error) {
	claims := Claims{
		UserClaims{
			ID:   id,
			Role: role,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert token to string")
	}

	return tokenString, nil
}
