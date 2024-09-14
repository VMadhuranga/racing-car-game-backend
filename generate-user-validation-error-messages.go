package main

import "github.com/go-playground/validator/v10"

func generateUserValidationErrorMessages(errors validator.ValidationErrors) userValidationErrors {
	messages := userValidationErrors{}

	for _, err := range errors {
		switch err.StructField() {
		case "Username":
			messages.Username = append(
				messages.Username,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		case "Password":
			messages.Password = append(
				messages.Password,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		}
	}

	return messages
}
