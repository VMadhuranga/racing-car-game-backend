package main

import (
	"log"
	"net/http"
	"strings"
)

func (api apiConfig) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println("Error Error getting bearer token: missing 'Bearer' prefix")
			respondWithError(w, 401, "Error getting bearer token")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if len(token) == 0 {
			log.Println("Error Error getting bearer token: missing token value")
			respondWithError(w, 401, "Error getting bearer token")
			return
		}

		_, err := validateJwt(token, api.accessTokenSecret)

		if err != nil {
			log.Printf("Error validating jwt: %s", err)
			respondWithError(w, 403, "Error validating jwt")
			return
		}

		next.ServeHTTP(w, r)
	})
}
