package main

import (
	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type apiConfig struct {
	queries            *database.Queries
	validate           *validator.Validate
	accessTokenSecret  string
	refreshTokenSecret string
}

type createUserPayload struct {
	Username        string `json:"username,omitempty" validate:"required,alpha"`
	Password        string `json:"password,omitempty" validate:"required,alphanum,min=5"`
	ConfirmPassword string `json:"confirm-password,omitempty"`
}

type userSignInPayload struct {
	Username string `json:"username,omitempty" validate:"required,alpha"`
	Password string `json:"password,omitempty" validate:"required,alphanum"`
}

type updateUsernamePayload struct {
	NewUsername string `json:"new-username,omitempty" validate:"required,alpha"`
}

type updatePasswordPayload struct {
	OldPassword        string `json:"old-password,omitempty" validate:"required,alphanum"`
	NewPassword        string `json:"new-password,omitempty" validate:"required,alphanum,min=5"`
	ConfirmNewPassword string `json:"confirm-new-password,omitempty"`
}

type updateBestTimePayLoad struct {
	BestTime string `json:"best-time,omitempty" validate:"required,numeric"`
}

type validationError struct {
	field, tag string
}

type userValidationErrorResponse struct {
	Username           []string `json:"username,omitempty"`
	Password           []string `json:"password,omitempty"`
	ConfirmPassword    []string `json:"confirm-password,omitempty"`
	NewUsername        []string `json:"new-username,omitempty"`
	OldPassword        []string `json:"old-password,omitempty"`
	NewPassword        []string `json:"new-password,omitempty"`
	ConfirmNewPassword []string `json:"confirm-new-password,omitempty"`
}

type userResponse struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	BestTime string    `json:"best-time,omitempty"`
}

type leaderBoardResponse struct {
	Id       uuid.UUID `json:"id,omitempty"`
	BestTime string    `json:"best-time,omitempty"`
	Username string    `json:"username,omitempty"`
}
