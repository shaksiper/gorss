package main

import (
	"fmt"
	"net/http"

	"github.com/shaksiper/go-tutorial/internal/auth"
	"github.com/shaksiper/go-tutorial/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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

		handler(writer, request, user)
	}
}
