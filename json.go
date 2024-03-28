package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Payload could not be marshaled %v", payload)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}

func respondWithError(writer http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println("Responding with 5xx error:", message)
	}
	type errResponse struct {
		Error string `json:"error"`
		/*
			{
				"error": "<message>"
			}
		*/
	}

	respondWithJson(writer, code, errResponse{
		Error: message,
	})
}
