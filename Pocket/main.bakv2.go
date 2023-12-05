package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the 'Frame.html' as the main page
		if r.URL.Path == "/" {
			data, err := os.ReadFile("../Templates/Frame.html")
			if err != nil {
				log.Printf("Error reading Frame.html: %v", err)
				http.Error(w, "we could not read file", http.StatusInternalServerError)
				return
			}
			w.Write(data)
		} else {
			// Handle requests for dynamic content
			// Trim the leading "/" and add the file extension if needed
			requestedFile := strings.TrimPrefix(r.URL.Path, "/")
			if !strings.HasSuffix(requestedFile, ".html") {
				requestedFile += ".html"
			}
			filePath := filepath.Join("../Templates", requestedFile)

			// Log the requested file path
			log.Printf("Requested file: %s", filePath)

			// Check if the file exists and serve it
			if data, err := os.ReadFile(filePath); err == nil {
				w.Write(data)
			} else {
				log.Printf("Error reading %s: %v", requestedFile, err)
				http.Error(w, "File not found.", http.StatusNotFound)
			}
		}
	})

	// Serve static files from the 'Static/css' directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Static"))))

	// Serve static files from the 'Static/css' directory
	//fs := http.FileServer(http.Dir("../Static"))
	//http.Handle("../Static/", http.StripPrefix("/Static/", fs))

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
