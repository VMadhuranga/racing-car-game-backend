package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api apiConfig) handleGetLeaderBoard(w http.ResponseWriter, r *http.Request) {
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

	leaderBoard, err := api.queries.GetLeaderBoard(r.Context())

	if err != nil {
		log.Printf("Error getting leader board: %s", err)
		respondWithError(w, 500, "Error getting leader board")
		return
	}

	lBRes := []leaderBoardResponse{}

	for _, lb := range leaderBoard {
		lBRes = append(lBRes, leaderBoardResponse{
			Id:       lb.ID,
			BestTime: lb.BestTime,
			Username: lb.Username,
		})
	}

	respondWithJson(w, 200, lBRes)
}
