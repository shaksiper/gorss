package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/shaksiper/go-tutorial/internal/database"
)

func (apiCfg *apiConfig) handlerFollowFeed(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, 412, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't follow feed (%s): %v", params.FeedID, err))
		return
	}

	respondWithJson(writer, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollowed(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollowed, err := apiCfg.DB.GetFeedFollowed(request.Context(), user.ID)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't follow feed for user (%s): %v", user.ID, err))
		return
	}

	respondWithJson(writer, 200, databaseFeedFollowsToFeedFollows(feedFollowed))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollowed(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollowedIDRaw := chi.URLParam(request, "feedFollowedID")
	feedFollowedID, err := uuid.Parse(feedFollowedIDRaw)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Could not parse feed ID (%s): %v", user.ID, err))
		return
	}
	// TODO: make this return deleated feed info
	err = apiCfg.DB.DeleteFeedFollowed(request.Context(), database.DeleteFeedFollowedParams{
		ID:     feedFollowedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Could not delete follow feed for user (%s): %v", user.ID, err))
		return
	}

	respondWithJson(writer, 200, struct{}{})
}
