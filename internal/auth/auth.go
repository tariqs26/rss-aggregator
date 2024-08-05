package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts the API key from the headers
// Example:
// Authorization: apiKey {api_key}
func GetApiKey(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")

	if authorization == "" {
		return "", errors.New("missing authorization header")
	}

	values := strings.Split(authorization, " ")

	if len(values) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if values[0] != "apiKey" {
		return "", errors.New("invalid authorization type")
	}

	return values[1], nil
}
