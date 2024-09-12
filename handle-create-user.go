package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (api apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
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
		errors := userValidationErrors{}

		for _, er := range err.(validator.ValidationErrors) {
			switch er.StructField() {
			case "Username":
				errors.Username = append(
					errors.Username,
					validationErrorMessages[validationError{er.StructField(), er.ActualTag()}],
				)
			case "Password":
				errors.Password = append(
					errors.Password,
					validationErrorMessages[validationError{er.StructField(), er.ActualTag()}],
				)
			}
		}

		respondWithValidationError(w, 400, errors)
		return
	}

	_, err = api.queries.GetUserByUsername(r.Context(), payload.Username)

	if err == nil {
		log.Printf("Could not validate payload: %s", "user exist")
		respondWithValidationError(w, 400, userValidationErrors{
			Username: []string{"User with this user name already exist"},
		})
		return
	}

	err = api.queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: payload.Username,
		Password: payload.Password,
	})

	if err != nil {
		log.Printf("Could not create user: %s", err)
		respondWithError(w, 500, "Could not create user")
		return
	}

	respondWithJson(w, 201, nil)
}
