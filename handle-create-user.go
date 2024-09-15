package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (api apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload userPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Printf("Error decoding payload: %s", err)
		respondWithError(w, 400, "Error decoding payload")
		return
	}

	err = api.validate.Struct(payload)

	if err != nil {
		log.Printf("Error validating payload: %s", err)

		respondWithValidationError(
			w,
			400,
			generateUserValidationErrorMessages(err.(validator.ValidationErrors)),
		)

		return
	}

	_, err = api.queries.GetUserByUsername(r.Context(), payload.Username)

	if err == nil {
		respondWithValidationError(w, 400, userValidationErrors{
			Username: []string{"User with this user name already exist"},
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		log.Printf("Error hashing password: %s", err)
		respondWithError(w, 500, "Error hashing password")
		return
	}

	err = api.queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: payload.Username,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Printf("Error creating user: %s", err)
		respondWithError(w, 500, "Error creating user")
		return
	}

	respondWithJson(w, 201, nil)
}
