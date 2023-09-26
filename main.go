package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID           int      `json:"id"`
	Locations    []string `json:"locations"`
	ConcertDates string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistDetails struct {
	Artist    Artist
	Locations Locations
	Dates     Dates
	Relations Relations
}

var tmpl = template.Must(template.ParseFiles("index.html"))
var artists []Artist

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Handle form data here
	w.Write([]byte("Form Submitted!"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Unable to get data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	artists = []Artist{}
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		http.Error(w, "Error decoding response body", http.StatusInternalServerError)
		return
	}

	var details []ArtistDetails
	for _, artist := range artists {
		locations := getLocations(artist.Locations)
		dates := getDates(artist.ConcertDates)
		relations := getRelations(artist.Relations)

		details = append(details, ArtistDetails{artist, locations, dates, relations})
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", details); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func getLocations(url string) Locations {
	var loc Locations
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error getting locations:", err)
		return loc
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&loc)
	return loc
}

func getDates(url string) Dates {
	var d Dates
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error getting dates:", err)
		return d
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&d)
	return d
}

func getRelations(url string) Relations {
	var rel Relations
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error getting relations:", err)
		return rel
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&rel)
	return rel
}
