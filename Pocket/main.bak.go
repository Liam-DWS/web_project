package main

import (
	"log"
	"net/http"
	"os"
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
