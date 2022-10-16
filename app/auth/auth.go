package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
