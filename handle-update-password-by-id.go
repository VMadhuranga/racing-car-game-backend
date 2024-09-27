package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (api apiConfig) handleUpdatePasswordById(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "userId"))

	if err != nil {
		log.Printf("Error parsing userId: %s", err)
		respondWithError(w, 404, "Error parsing userId")
		return
	}

	user, err := api.queries.GetUserById(r.Context(), userId)

	if err != nil {
		log.Printf("Error getting user: %s", err)
		respondWithError(w, 404, "Error getting user")
		return
	}

	payload := updatePasswordPayload{}
	err = json.NewDecoder(r.Body).Decode(&payload)

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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.OldPassword))

	if err != nil {
		log.Printf("Error comparing passwords: %s", err)

		respondWithValidationError(w, 401, userValidationErrorResponse{
			OldPassword: []string{"Incorrect old password"},
		})

		return
	}

	if payload.NewPassword != payload.ConfirmNewPassword {
		respondWithValidationError(w, 400, userValidationErrorResponse{
			ConfirmNewPassword: []string{"New passwords do not match"},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), 10)

	if err != nil {
		log.Printf("Error hashing password: %s", err)
		respondWithError(w, 500, "Error hashing password")
		return
	}

	err = api.queries.UpdatePasswordById(r.Context(), database.UpdatePasswordByIdParams{
		ID:       userId,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Printf("Error updating username: %s", err)
		respondWithError(w, 500, "Error updating username")
		return
	}

	respondWithJson(w, 204, nil)
}
