package main

import (
    "bytes"
    "net/http"
    "testing"
)

func TestLoginAPI(t *testing.T) {
    // Prepare request payload
    payload := `{"username": "beaconuser", "password": "beaconpassword"}`
    req, err := http.NewRequest("POST", "http://localhost:8081/login", bytes.NewBufferString(payload))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Cookie", "mysession=MTcxMjIwODk3MXxEWDhFQVFMX2dBQUJFQUVRQUFBb180QUFBUVp6ZEhKcGJtTURBQUtZbVZoWTI5dWRYTmxjZz09fKhxzQEdk94NTdYSB30r_iDozWGvNW3UpKbMgVKUzJY3")

    // Perform the request
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
}

