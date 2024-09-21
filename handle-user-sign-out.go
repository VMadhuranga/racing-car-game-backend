package main

import (
	"net/http"
)

func (api apiConfig) handleUserSignOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:        "jwt",
		Value:       "",
		HttpOnly:    true,
		Secure:      true,
		SameSite:    http.SameSiteNoneMode,
		Partitioned: true,
		MaxAge:      -1,
	}

	http.SetCookie(w, &cookie)
	respondWithJson(w, 204, nil)
}
