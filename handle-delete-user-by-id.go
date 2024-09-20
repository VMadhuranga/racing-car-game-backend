package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api apiConfig) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
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

	err = api.queries.DeleteUserById(r.Context(), userId)

	if err != nil {
		log.Printf("Error deleting user: %s", err)
		respondWithError(w, 500, "Error deleting user")
		return
	}

	respondWithJson(w, 204, nil)
}
