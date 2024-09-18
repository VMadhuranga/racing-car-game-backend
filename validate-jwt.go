package main

import "github.com/golang-jwt/jwt/v5"

func validateJwt(token, jwtSecret string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	sub, err := t.Claims.GetSubject()

	if err != nil {
		return "", err
	}

	return sub, nil
}
