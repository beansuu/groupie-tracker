package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Artist struct
type Artist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Year           int    `json:"year"`
	ImageURL       string `json:"image_url"`
	FirstAlbumDate string `json:"first_album_date"`
}

type Location struct {
	ID              string `json:"id"`
	LastConcert     string `json:"last_concert"`
	UpcomingConcert string `json:"upcoming_concert"`
}

type Date struct {
	ID              string `json:"id"`
	LastConcert     string `json:"last_concert"`
	UpcomingConcert string `json:"upcoming_concert"`
}

type Relation struct {
	ArtistID   string `json:"artist_id"`
	LocationID string `json:"location_id"`
	DateID     string `json:"date_id"`
}

// artists slice
var artists = []Artist{
	{
		ID:             "1",
		Name:           "Band 1",
		Year:           2000,
		ImageURL:       "http://example.com/image1.jpg",
		FirstAlbumDate: "2001-01-01",
	},
}

var locations = []Location{
	{
		ID:              "1",
		LastConcert:     "New York",
		UpcomingConcert: "Los Angeles",
	},
}

var dates = []Date{
	{
		ID:              "1",
		LastConcert:     "2022-09-15",
		UpcomingConcert: "2023-10-20",
	},
}

var relations = []Relation{
	{
		ArtistID:   "1",
		LocationID: "1",
		DateID:     "1",
	},
}

func main() {
	// Serve static assets like CSS
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// API route for artists
	http.HandleFunc("/artists", artistsHandler)

	// API route for locations
	http.HandleFunc("/locations", locationsHandler)

	// Route for home page
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/dates", datesHandler)

	http.HandleFunc("/relations", relationsHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Check for correct method
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Encode and send the artists data
	if err := json.NewEncoder(w).Encode(artists); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html file
	http.ServeFile(w, r, "index.html")
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func datesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func relationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}
