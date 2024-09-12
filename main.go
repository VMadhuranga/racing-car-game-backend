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
		log.Fatalf("Could not load .env file: %s", err)
	}

	dbUri := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", dbUri)

	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}

	api := apiConfig{
		queries:  database.New(db),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_BASE_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	v1router.Post("/users", api.handleCreateUser)

	router.Mount("/v1", v1router)
	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port: %s", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Could not listen on server: %s", err)
	}
}
