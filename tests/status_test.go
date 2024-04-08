package main

import (
    "io/ioutil"
    "net/http"
    "testing"
)

func TestStatusAPI(t *testing.T) {
    // Prepare request
    req, err := http.NewRequest("GET", "http://localhost:8081/status", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // Send request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        t.Fatalf("Failed to send request: %v", err)
    }
    defer resp.Body.Close()

    // Check the response status code
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }

    // Read the response body
    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    // Optionally, add assertions for the response body if needed
}

