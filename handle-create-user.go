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
	var payload createUserPayload
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
		respondWithValidationError(w, 400, userValidationErrorResponse{
			Username: []string{"User with this user name already exist"},
		})

		return
	}

	if payload.Password != payload.ConfirmPassword {
		respondWithValidationError(w, 400, userValidationErrorResponse{
			ConfirmPassword: []string{"Passwords do not match"},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		log.Printf("Error hashing password: %s", err)
		respondWithError(w, 500, "Error hashing password")
		return
	}

	userId := uuid.New()

	err = api.queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       userId,
		Username: payload.Username,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Printf("Error creating user: %s", err)
		respondWithError(w, 500, "Error creating user")
		return
	}

	err = api.queries.AddUserToLeaderBoard(r.Context(), database.AddUserToLeaderBoardParams{
		ID:       uuid.New(),
		BestTime: "0.0",
		UserID:   userId,
	})

	if err != nil {
		log.Printf("Error adding user to leader board: %s", err)
		respondWithError(w, 500, "Error adding user to leader board")
		return
	}

	respondWithJson(w, 201, nil)
}
