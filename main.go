package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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

type LocationsData struct {
	Index []Locations `json:"index"`
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
	Artist   Artist
	Members  string
	Concerts string
}

type Item struct {
	ID       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}

type ConcertData struct {
	Index []Item `json:"index"`
}

type TemplateData struct {
	Artists     []ArtistDetails
	Suggestions []string
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

	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/relation")
	concertData := ConcertData{}
	if err := json.NewDecoder(resp.Body).Decode(&concertData); err != nil {
		http.Error(w, "Error decoding response body", http.StatusInternalServerError)
		return
	}

	var details []ArtistDetails
	for _, artist := range artists {
		members := strings.Join(artist.Members, ", ")
		concerts := parseConcerts(concertData.Index[artist.ID-1].Concerts)

		details = append(details, ArtistDetails{artist, members, concerts})
	}
	data := TemplateData{
		Artists:     details,
		Suggestions: getSuggestions(artists),
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func parseConcerts(concerts map[string][]string) string {
	str := ""
	for location, dates := range concerts {
		for _, date := range dates {
			formatted := fmt.Sprintf("%s: %s", date, formatLocation(location))
			str += formatted + "\n"
		}
	}
	return str
}

func getSuggestions(artists []Artist) []string {
	suggestions := make(map[string]int)
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	locationData := LocationsData{}
	if err := json.NewDecoder(resp.Body).Decode(&locationData); err == nil {
		for _, location := range locationData.Index {
			for _, loc := range location.Locations {
				formatted := formatLocation(loc)
				suggestions[formatted] = 1
			}

		}
	}
	for _, artist := range artists {
		suggestions[artist.Name] = 1
		for _, member := range artist.Members {
			suggestions[member] = 1
		}
	}

	fmt.Println(suggestions)
	var suggs []string
	for k := range suggestions {
		suggs = append(suggs, k)
	}
	return suggs
}

func formatLocation(loc string) string {
	parts := strings.Split(loc, "-")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	formatted := strings.Join(parts, ", ")
	return formatted
}
