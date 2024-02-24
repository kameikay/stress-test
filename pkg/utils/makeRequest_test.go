package utils

import (
	"net/http"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	// Mock server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	go http.ListenAndServe(":8080", nil)

	// Test MakeRequest function
	resp, err := MakeRequest("http://localhost:8080")
	if err != nil {
		t.Errorf("MakeRequest failed with error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got: %s", resp.Status)
	}
}
