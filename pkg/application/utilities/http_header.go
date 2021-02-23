package utilities

import (
	"net/http"
	"strings"
)

// ExtractBearerToken extract token from header authorization Bearer
func ExtractBearerToken(r *http.Request) string {
	val := strings.TrimSpace(r.Header.Get("authorization"))

	if val == "" {
		return ""
	}

	splitToken := strings.Split(val, " ")

	if len(splitToken) < 2 {
		return ""
	}

	return splitToken[1]
}
