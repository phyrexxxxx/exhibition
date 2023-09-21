package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey() retrieves the API key from the "Authorization" header in the request headers
// Authorization: ApiKey <key>
func GetApiKey(headers http.Header) (string, error) {
	authz := headers.Get("Authorization")
	if authz == "" {
		return "", errors.New("no authorization header found")
	}

	tokens := strings.Split(authz, " ")
	if len(tokens) != 2 {
		return "", errors.New("invalid authorization header")
	}
	if tokens[0] != "ApiKey" {
		return "", errors.New("usage: Authorization: ApiKey <key>")
	}

	return tokens[1], nil
}
