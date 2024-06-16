package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/go-tutorial/internal/auth"
	"github.com/shaksiper/go-tutorial/internal/database"
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
func (apiCfg *apiConfig) handlerGetUserByAPIKey(writer http.ResponseWriter, request *http.Request) {
	apiKey, err := auth.GetAPIKey(request.Header)
	if err != nil {
		respondWithError(writer, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(request.Context(), apiKey)
	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Could not fetch the user: %v", err))
        return
	}

	respondWithJson(writer, 200, databaseUserToUser(user))
}
