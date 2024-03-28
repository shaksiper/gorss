package main

import "net/http"

func handlerErr(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, 400, "Something went wrong")
}
