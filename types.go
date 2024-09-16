package main

import (
	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

type apiConfig struct {
	queries            *database.Queries
	validate           *validator.Validate
	accessTokenSecret  string
	refreshTokenSecret string
}

type userPayload struct {
	Username        string `validate:"required"`
	Password        string `validate:"required,alphanum"`
	ConfirmPassword string `json:",omitempty"`
}

type validationError struct {
	field, tag string
}

type userValidationErrorResponse struct {
	Username        []string `json:"username,omitempty"`
	Password        []string `json:"password,omitempty"`
	ConfirmPassword []string `json:"confirm-password,omitempty"`
}
