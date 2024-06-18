package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/go-tutorial/internal/database"
)

func (apicfg *apiConfig) handlerCreateFeed(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, 412, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apicfg.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't create feed (%s): %v", params.Name, err))
		return
	}

	respondWithJson(writer, 201, feed)
}

func (apicfg *apiConfig) handlerGetFeeds(writer http.ResponseWriter, request *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(request.Context())
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Could not retrieve feeds: %v", err))
		return
	}

	respondWithJson(writer, 200, databaseFeedsToFeeds(feeds))
}
