package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example: Authorization: ApiKey {ApiKey}
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}

	pair := strings.Split(authHeader, " ")
	if len(pair) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if pair[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return pair[1], nil
}
