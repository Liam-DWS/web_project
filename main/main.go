package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filePath, contentType, suffix string

		if r.URL.Path == "/" {
			filePath = "../Templates/Frame.html"
			contentType = "text/html"
		} else if strings.HasPrefix(r.URL.Path, "/static/") {
			// Serving a static file
			suffix = ".css"
			filePath = "../Static" + strings.TrimPrefix(r.URL.Path, "/static")
			contentType = "text/css" // Assuming all static files are CSS
		} //else {
		// Serving a template HTML file
		//suffix = ".html"
		//filePath = "../Templates" + r.URL.Path + suffix
		//contentType = "text/html"
		//}

		serveFile(w, r, filePath, contentType)
	})

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
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
