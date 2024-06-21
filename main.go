package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// router := chi.NewRouter()

	// router.Use(
	// 	cors.Handler(cors.Options{
	// 		AllowedOrigins:   []string{"https://*", "http://*"},
	// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 		AllowedHeaders:   []string{"*"},
	// 		ExposedHeaders:   []string{"Link"},
	// 		AllowCredentials: false,
	// 		MaxAge:           300,
	// 	}),
	// )

	// router := chi.NewRouter()
	router := http.NewServeMux()
	router.HandleFunc("/ready", handlerReadiness)
	router.HandleFunc("/error", handlerErr)

	router.HandleFunc("POST /user", apiCfg.handlerCreateUser)
	router.HandleFunc("GET /user", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))

	router.HandleFunc("POST /feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	router.HandleFunc("GET /feed", apiCfg.handlerGetFeeds)

	router.HandleFunc("POST /feed_follow", apiCfg.middlewareAuth(apiCfg.handlerFollowFeed))
	router.HandleFunc("GET /feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollowed))
	router.HandleFunc("DELETE /feed_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollowed))

	router.HandleFunc("GET /post", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1Router := http.NewServeMux()
	v1Router.Handle("/v1/", http.StripPrefix("/v1", router))
	// router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: corsMiddleware(v1Router),
		Addr:    ":" + port,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start the server")
		log.Fatal(err)
	}

	fmt.Println("Hello World")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("AllowedOrigins", "https://* http://*")
		w.Header().Set("AllowedMethods", "POST PUT DELETE OPTIONS")
		w.Header().Set("AllowedHeaders", "*")
		w.Header().Set("ExposedHeaders", "Link")
		w.Header().Set("AllowCredentials", "false")
		w.Header().Set("MaxAge", "300")
		next.ServeHTTP(w, r)
	})
}
