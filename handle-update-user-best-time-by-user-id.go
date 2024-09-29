package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api apiConfig) handleUpdateUserBestTimeByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "userId"))

	if err != nil {
		log.Printf("Error parsing userId: %s", err)
		respondWithError(w, 404, "Error parsing userId")
		return
	}

	_, err = api.queries.GetUserById(r.Context(), userId)

	if err != nil {
		log.Printf("Error getting user: %s", err)
		respondWithError(w, 404, "Error getting user")
		return
	}

	payload := updateBestTimePayLoad{}
	err = json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Printf("Error decoding payload: %s", err)
		respondWithError(w, 400, "Error decoding payload")
		return
	}

	err = api.validate.Struct(payload)

	if err != nil {
		log.Printf("Error validating payload: %s", err)
		respondWithError(w, 400, "Error validating payload")
		return
	}

	err = api.queries.UpdateUserBestTimeByUserId(r.Context(), database.UpdateUserBestTimeByUserIdParams{
		UserID:   userId,
		BestTime: payload.BestTime,
	})

	if err != nil {
		log.Printf("Error updating user best time: %s", err)
		respondWithError(w, 500, "Error updating user best time")
		return
	}

	respondWithJson(w, 204, nil)
}
