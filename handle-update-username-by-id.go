package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (api apiConfig) handleUpdateUsernameById(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "userId"))

	if err != nil {
		log.Printf("Error parsing userId: %s", err)
		respondWithError(w, 400, "Error parsing userId")
		return
	}

	_, err = api.queries.GetUserById(r.Context(), userId)

	if err != nil {
		log.Printf("Error getting user: %s", err)
		respondWithError(w, 404, "Error getting user")
		return
	}

	payload := updateUsernamePayload{}
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

	user, err := api.queries.GetUserByUsername(r.Context(), payload.NewUsername)

	if err == nil && user.ID != userId {
		respondWithValidationError(w, 400, userValidationErrorResponse{
			NewUsername: []string{"User with this user name already exist"},
		})

		return
	}

	err = api.queries.UpdateUsernameById(r.Context(), database.UpdateUsernameByIdParams{
		ID:       userId,
		Username: payload.NewUsername,
	})

	if err != nil {
		log.Printf("Error updating username: %s", err)
		respondWithError(w, 500, "Error updating username")
		return
	}

	respondWithJson(w, 204, nil)
}
