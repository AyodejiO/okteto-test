package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ExpectedResponse struct {
	Success bool `json:"success"`
	Data  struct {
		Message   string `json:"message"`
	} `json:"data"`
}

func TestGetIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIndexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
	}

	var er ExpectedResponse
	if err = json.Unmarshal([]byte(rr.Body.Bytes()), &er); err != nil {
		t.Fatal(err)
	}

	if er.Success != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Success, true)
	}

	if er.Data.Message != "Welcome to the Okteto API" {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Data.Message, "Welcome to the Okteto API")
	}
}

func TestMissingRouteHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/not-found", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Custom404Handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusNotFound)
	}

	var er ExpectedResponse
	if err = json.Unmarshal([]byte(rr.Body.Bytes()), &er); err != nil {
		t.Fatal(err)
	}

	if er.Success != false {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Success, false)
	}

	if er.Data.Message != "Requested resource not found" {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Data.Message, "Not Found")
	}
}