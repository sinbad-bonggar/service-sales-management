package utilities

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()

	data := "This is the response data"

	meta := "This is the response meta"

	JSON(w, r, nil, http.StatusOK, nil, data, meta)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":200,"type":null,"data":"This is the response data","message":null,"meta":"This is the response meta","code":null}`

	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestJSONSuccess(t *testing.T) {
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()

	data := "This is the response data"

	meta := "This is the response meta"

	JSONSuccess(w, r, data, meta)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":200,"type":null,"data":"This is the response data","message":null,"meta":"This is the response meta","code":null}`

	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestJSONError(t *testing.T) {
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()

	message := "This is the response message"

	status := 400

	resType := "This is the response type"

	JSONError(w, r, message, status, resType)

	if status := w.Code; status != status {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":400,"type":"This is the response type","data":null,"message":"This is the response message","meta":null,"code":null}`

	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
