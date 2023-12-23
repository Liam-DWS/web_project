// Package main contains the entry point and HTTP server implementation
// for the web application.
package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//serves the filepathe for frame.html
		data, err := os.ReadFile("../Templates/Frame.html")
		if err != nil {
			http.Error(w, "we could not read file", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func testServer(t *testing.T) {
	t.Run("starts HTTP server", func(t *testing.T) {
		// Create a test server using httptest.NewServer()
		ts := httptest.NewServer(http.HandlerFunc(main)) // Pass main as the handler
		defer ts.Close()

		resp, err := http.Get(ts.URL) // Use the test server's URL
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("serves index.html", func(t *testing.T) {
		// Create a request and recorder
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		// Call the handler directly, passing request and recorder
		main(w, req) // Pass both w and req to main

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)
		}

		body := w.Body.String()
		if !strings.Contains(body, "<html>") {
			t.Errorf("got body %q, want to contain html", body)
		}
	})
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string, contentType string) {
	// Check if the file exists
	if _, err := os.Stat(filePath); err != nil {
		log.Printf("File not found: %s, error: %v", filePath, err)
		http.NotFound(w, r)
		return
	}

	// Set the content type and serve the file
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, filePath)
}
