package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (api apiConfig) handleRefresh(w http.ResponseWriter, r *http.Request) {
	jwtCookie, err := r.Cookie("jwt")

	if err != nil {
		log.Printf("Error getting jwt cookie: %s", err)
		respondWithError(w, 401, "Error getting jwt cookie")
		return
	}

	tokenSubject, err := validateJwt(jwtCookie.Value, api.refreshTokenSecret)

	if err != nil {
		log.Printf("Error validating jwt: %s", err)
		respondWithError(w, 403, "Error validating jwt")
		return
	}

	userId, err := uuid.Parse(tokenSubject)

	if err != nil {
		log.Printf("Error parsing token subject: %s", err)
		respondWithError(w, 400, "Error parsing token subject")
		return
	}

	user, err := api.queries.GetUserById(r.Context(), userId)

	if err != nil {
		log.Printf("Error getting user: %s", err)
		respondWithError(w, 404, "Error getting user")
		return
	}

	accessToken, err := createJwt(time.Minute, user.ID.String(), api.accessTokenSecret)

	if err != nil {
		log.Printf("Error creating access token: %s", err)
		respondWithError(w, 500, "Error creating access token")
		return
	}

	refreshToken, err := createJwt(24*time.Hour, user.ID.String(), api.refreshTokenSecret)

	if err != nil {
		log.Printf("Error creating refresh token: %s", err)
		respondWithError(w, 500, "Error creating refresh token")
		return
	}

	cookie := http.Cookie{
		Name:        "jwt",
		Value:       refreshToken,
		HttpOnly:    true,
		Secure:      true,
		SameSite:    http.SameSiteNoneMode,
		Partitioned: true,
		MaxAge:      1 * 24 * 60 * 60 * 1000,
	}

	http.SetCookie(w, &cookie)

	respondWithJson(
		w,
		200,
		map[string]string{"userId": user.ID.String(), "accessToken": accessToken},
	)
}
