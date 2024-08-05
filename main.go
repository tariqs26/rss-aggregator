package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/scraper"
)

type ApiConfig struct {
	DB *database.Queries
}

var apiConfig ApiConfig

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	apiConfig.DB = database.New(db)

	go scraper.StartScraping(apiConfig.DB, 10, time.Minute)

	router := chi.NewRouter()

	corsOptions := cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	router.Use(cors.Handler(corsOptions))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/v1", v1Router())

	log.Printf("Server is running on http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(err)
	}
}
