package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/gorss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, 412, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Couldn't create user (%s): %v", params.Name, err))
		return
	}

	respondWithJson(writer, 201, user)
}

// AUthenticated endpoint to retrieve user information
func (apiCfg *apiConfig) handlerGetUserByAPIKey(writer http.ResponseWriter, request *http.Request, user database.User) {
	respondWithJson(writer, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(writer http.ResponseWriter, request *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(request.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Could not retrieve posts for user %v: %v", user.ID, err))
		return
	}

	respondWithJson(writer, 200, databasePostsToPosts(posts))
}
