package main

import (
	"auth-rate-limit-api/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {

	// setting up request and response recorder
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recordResponse := httptest.NewRecorder()

	// Calling the home handler
	auth.HomePage(recordResponse, request)

	// Inspecting the results
	result := recordResponse.Result()

	// defer closing the response body
	defer result.Body.Close()

	if result.StatusCode != http.StatusBadGateway {
		t.Errorf("Home Handler not returning the correct status code %d", result.StatusCode)
	}
}
