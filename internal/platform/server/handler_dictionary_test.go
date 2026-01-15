package server

import (
	"encoding/json"
	"impostor/internal/domain"
	"net/http/httptest"
	"testing"
)

func TestGetRandomWordHandler(t *testing.T) {
	// Initialize Server
	s := NewServer()

	// Test Case 1: Valid Request
	req := httptest.NewRequest("GET", "/api/word?category=General&lang=en", nil)
	resp, err := s.App.Test(req)
	if err != nil {
		t.Fatalf("App.Test error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var word domain.WordPair
	if err := json.NewDecoder(resp.Body).Decode(&word); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if word.Real == "" || word.Trap == "" {
		t.Errorf("Expected valid word pair, got empty values: %v", word)
	}

	// Test Case 2: Spanish
	reqES := httptest.NewRequest("GET", "/api/word?category=General&lang=es", nil)
	respES, _ := s.App.Test(reqES)
	if respES.StatusCode != 200 {
		t.Errorf("Expected status 200 for ES, got %d", respES.StatusCode)
	}
}

func TestGetCategoriesHandler(t *testing.T) {
	s := NewServer()

	req := httptest.NewRequest("GET", "/api/categories?lang=en", nil)
	resp, err := s.App.Test(req)
	if err != nil {
		t.Fatalf("App.Test error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var result map[string][]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(result["categories"]) == 0 {
		t.Error("Expected categories list, got empty")
	}
}
