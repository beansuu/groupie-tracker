package main

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	ImageURL string `json:"image_url"`
}

func main() {
	http.HandleFunc("/artists", ArtistsHandler)
	http.ListenAndServe(":8080", nil)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists := []Artist{
		{Name: "Band 1", Year: 2000, ImageURL: "http://example.com/image1.jpg"},
		// ... add other artists ...
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
