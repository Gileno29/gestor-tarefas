package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	// Arrange
	users := []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}))

	// Act
	response, err := http.Get(server.URL + "/users")

	// Assert
	if err != nil {
		t.Errorf("Error getting users: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	var actualUsers []User
	err = json.NewDecoder(response.Body).Decode(&actualUsers)

	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	if !reflect.DeepEqual(actualUsers, users) {
		t.Errorf("Expected users %v, got %v", users, actualUsers)
	}
}
