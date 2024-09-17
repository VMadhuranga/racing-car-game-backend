package main

import "github.com/go-playground/validator/v10"

func generateUserValidationErrorMessages(errors validator.ValidationErrors) userValidationErrorResponse {
	messages := userValidationErrorResponse{}

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
		case "NewUsername":
			messages.NewUsername = append(
				messages.NewUsername,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		case "OldPassword":
			messages.OldPassword = append(
				messages.OldPassword,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		case "NewPassword":
			messages.NewPassword = append(
				messages.NewPassword,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		case "ConfirmNewPassword":
			messages.ConfirmNewPassword = append(
				messages.ConfirmNewPassword,
				validationErrorMessages[validationError{err.StructField(), err.ActualTag()}],
			)
		}
	}

	return messages
}
