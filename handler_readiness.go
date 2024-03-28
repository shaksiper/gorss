package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJson(writer, 200, struct{}{})
}
