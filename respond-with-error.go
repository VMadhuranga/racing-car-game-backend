package main

import "net/http"

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	respondWithJson(w, statusCode, map[string]string{"error": errorMessage})
}
