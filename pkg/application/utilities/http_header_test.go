package utilities

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractBearerToken(t *testing.T) {
	const tokenValue = "abcd"

	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	// ------------
	value := ExtractBearerToken(r)
	if value != "" {
		t.Errorf("ExtractBearerToken not return empty string")
	}
	// ------------

	// ------------
	r.Header.Set("authorization", tokenValue)

	value = ExtractBearerToken(r)

	if value != "" {
		t.Errorf("ExtractBearerToken not return empty string")
	}
	// ------------

	// ------------
	r.Header.Set("authorization", fmt.Sprintf("Bearer %s", tokenValue))

	value = ExtractBearerToken(r)

	if value != tokenValue {
		t.Errorf("ExtractBearerToken not return valid value")
	}
	// ------------
}
