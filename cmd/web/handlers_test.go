package main

import (
	"kkanat/forum/internal/assert"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	t.Parallel()

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

// func TestPostCreate(t *testing.T) {
// 	// Initialize a new test application
// 	app := newTestApplication(t)

// 	// Initialize a new test server with the postCreate handler
// 	ts := newTestServer(t, http.HandlerFunc(app.postCreate))

// 	// Test Case 1: Successful post creation
// 	t.Run("successful_post_creation", func(t *testing.T) {
// 		// Create a new POST request with form data
// 		formData := "title=Test+Title&description=Test+Description&category=TestTag"
// 		reqBody := strings.NewReader(formData)
// 		req := httptest.NewRequest(http.MethodPost, ts.URL+"/post/create", reqBody)
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 		// Perform the request
// 		resp, err := ts.Client().Do(req)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer resp.Body.Close()

// 		// Check if the response status code is as expected
// 		if resp.StatusCode != http.StatusSeeOther {
// 			t.Errorf("expected status %d, got %d", http.StatusSeeOther, resp.StatusCode)
// 		}
// 	})

// 	// Test Case 2: Invalid request method
// 	t.Run("invalid_request_method", func(t *testing.T) {
// 		// Create a new GET request instead of POST
// 		req := httptest.NewRequest(http.MethodGet, ts.URL+"/post/create", nil)

// 		// Perform the request
// 		resp, err := ts.Client().Do(req)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer resp.Body.Close()

// 		// Check if the response status code is as expected
// 		if resp.StatusCode != http.StatusMethodNotAllowed {
// 			t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
// 		}
// 	})

// 	// Add more test cases as needed...
// }
