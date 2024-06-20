package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/shaksiper/gorss/internal/database"

	_ "github.com/lib/pq" // postgres driver for sqlc
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// read .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not specified")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Connection cannot be empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Database connection could not be established")
	}

	db := database.New(conn)
	queries := db

	apiCfg := apiConfig{
		DB: queries,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handlerReadiness)
	v1Router.HandleFunc("/error", handlerErr)

	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))

	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed", apiCfg.handlerGetFeeds)

	v1Router.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerFollowFeed))
	v1Router.Get("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollowed))
	v1Router.Delete("/feed_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollowed))

	v1Router.Get("/post", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start the server")
		log.Fatal(err)
	}

	fmt.Println("Hello World")
}
