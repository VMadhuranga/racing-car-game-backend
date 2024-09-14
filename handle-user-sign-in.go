package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (api apiConfig) handleUserSignIn(w http.ResponseWriter, r *http.Request) {
	var payload userPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Printf("Could not decode payload: %s", err)
		respondWithError(w, 400, "Could not decode payload")
		return
	}

	err = api.validate.Struct(payload)

	if err != nil {
		log.Printf("Could not validate payload: %s", err)

		respondWithValidationError(
			w,
			400,
			generateUserValidationErrorMessages(err.(validator.ValidationErrors)),
		)

		return
	}

	user, err := api.queries.GetUserByUsername(r.Context(), payload.Username)

	if err != nil {
		log.Printf("Could not validate payload: %s", err)

		respondWithValidationError(w, 401, userValidationErrors{
			Username: []string{"Incorrect username"},
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		log.Printf("Could not validate payload: %s", err)

		respondWithValidationError(w, 401, userValidationErrors{
			Password: []string{"Incorrect password"},
		})

		return
	}

	accessToken, err := createJwt(time.Minute, user.ID.String(), api.accessTokenSecret)

	if err != nil {
		log.Printf("Could not create access token: %s", err)
		respondWithError(w, 500, "Could not create access token")
		return
	}

	refreshToken, err := createJwt(24*time.Hour, user.ID.String(), api.refreshTokenSecret)

	if err != nil {
		log.Printf("Could not create refresh token: %s", err)
		respondWithError(w, 500, "Could not create refresh token")
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
