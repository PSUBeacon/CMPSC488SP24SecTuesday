package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppliancesEndpoint(t *testing.T) {
	// Create a request body with login credentials
	loginData := map[string]string{
		"username": "beaconuser",
		"password": "beaconpassword",
	}
	loginBody, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Failed to marshal login request body: %v", err)
	}

	// Make a POST request to the /login endpoint on localhost to obtain the JWT token
	loginResp, err := http.Post("http://localhost:8081/login", "application/json", bytes.NewBuffer(loginBody))
	if err != nil {
		t.Fatalf("Failed to make login request: %v", err)
	}
	defer loginResp.Body.Close()

	// Check if login was successful
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	// Decode the login response body to extract the JWT token
	var loginResponseBody map[string]string
	err = json.NewDecoder(loginResp.Body).Decode(&loginResponseBody)
	if err != nil {
		t.Fatalf("Failed to decode login response body: %v", err)
	}

	jwtToken, ok := loginResponseBody["token"]
	if !ok {
		t.Fatalf("JWT token not found in login response")
	}

	// Create a request body with test data for the /appliances endpoint
	appliancesData := map[string]interface{}{
		// Add your test data for the appliances endpoint
	}
	appliancesBody, err := json.Marshal(appliancesData)
	if err != nil {
		t.Fatalf("Failed to marshal appliances request body: %v", err)
	}

	// Create a new HTTP request with the JWT token in the Authorization header
	req, err := http.NewRequest("POST", "http://localhost:8081/appliances", bytes.NewBuffer(appliancesBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Content-Type", "application/json")

	// Make a request to the /appliances endpoint with the JWT token
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request to /appliances endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the response body
	var responseBody interface{} // Change the type based on your expected response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Add your assertions for the response body here
}

