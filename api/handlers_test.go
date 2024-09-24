package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApi_CreateUser(t *testing.T) {

	// Create a new HTTP POST request
	requestBody := []byte(`{"username":"go_tester","password":"go_pass"}`)
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function to handle the request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here you can handle the request and write the response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	})

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestApi_LoginUser(t *testing.T) {
	// Create a new HTTP GET request with query parameters
	params := url.Values{}
	params.Add("username", "testuser")
	params.Add("password", "testpass")
	req, err := http.NewRequest("GET", "/login?"+params.Encode(), nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function to handle the request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here you can handle the request and write the response
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		if username == "testuser" && password == "testpass" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"success"}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status":"unauthorized"}`))
		}
	})

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"status":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
