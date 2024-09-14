package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJwt(expTime time.Duration, userId, jwtSecret string) (string, error) {
	currentTime := time.Now()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "RCG",
		IssuedAt:  jwt.NewNumericDate(currentTime),
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(expTime)),
		Subject:   userId,
	}).SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}
