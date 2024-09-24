package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/VMadhuranga/racing-car-game-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbUri := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", dbUri)

	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	api := apiConfig{
		queries:            database.New(db),
		validate:           validator.New(validator.WithRequiredStructEnabled()),
		accessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		refreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_BASE_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	// public routes
	v1router.Group(func(r chi.Router) {
		r.Post("/sign-in", api.handleUserSignIn)
		r.Get("/sign-out", api.handleUserSignOut)
		r.Get("/refresh", api.handleRefresh)
		r.Post("/users", api.handleCreateUser)
	})

	// private routes
	v1router.Group(func(r chi.Router) {
		r.Use(api.authenticate)
		r.Get("/users/{userId}", api.handleGetUserById)
		r.Delete("/users/{userId}", api.handleDeleteUserById)
		r.Patch("/users/{userId}/username", api.handleUpdateUsernameById)
		r.Patch("/users/{userId}/password", api.handleUpdatePasswordById)
		r.Patch("/users/{userId}/leader-board/best-time", api.updateUserBestTimeByUserId)
	})

	router.Mount("/v1", v1router)
	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port: %s", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Error listening on server: %s", err)
	}
}
