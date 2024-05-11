package auth

import (
	"errors"
	"net/http"
	"strings"
)

/**
*	GetAPIKey extracts the API key from the headers of an HTTP request
 */

/**
*	Example:
*		Authorization: ApiKey {api_key}
 */
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization");
	if val == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentication information")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authentication method")
	}

	return vals[1], nil
}