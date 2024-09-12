package main

import (
	"net/http"
)

func respondWithValidationError(w http.ResponseWriter, statusCode int, errors interface{}) {
	respondWithJson(w, statusCode, map[string]interface{}{"errors": errors})
}
