package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupHandler(t *testing.T) {
	// Create a request body with test data
	signupData := map[string]string{
		"firstname": "John",
		"lastname":  "Doe",
		"password":  "password123",
		"username":  "johndoe",
	}
	requestBody, err := json.Marshal(signupData)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Make a POST request to the /signup endpoint on localhost
	resp, err := http.Post("http://localhost:8081/signup", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the response body
	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check the response message
	expectedMessage := "User created succesfully"
	assert.Equal(t, expectedMessage, responseBody["message"])
}

