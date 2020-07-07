package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// API link to JSON of Studio Ghibli films
const API = "https://ghibliapi.herokuapp.com/films"

func main() {

	filmPtr := flag.String("film", "Castle in the Sky", "A Studio Ghibli Film\nReturns the film of interest's description, director, producer, release date, and Rotten Tomato score.")

	flag.Parse()

	if *filmPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	contents := getResponseData(API)

	var films []Film
	json.Unmarshal(contents, &films)

	for _, film := range films {
		if film.Title == *filmPtr {
			fmt.Println("")
			fmt.Printf("Title: %s\n", film.Title)
			fmt.Println("------------------------")
			fmt.Printf("Description: %s\n", film.Description)
			fmt.Println("------------------------")
			fmt.Printf("Director: %s\n", film.Director)
			fmt.Println("------------------------")
			fmt.Printf("Producer: %s\n", film.Producer)
			fmt.Println("------------------------")
			fmt.Printf("Release Year: %s\n", film.ReleaseDate)
			fmt.Println("------------------------")
			fmt.Printf("Rotten Tomato Score: %s\n", film.RtScore)
			fmt.Println("------------------------")
			fmt.Println("")
		}
	}
}

//Film data type
type Film struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Director    string   `json:"director"`
	Producer    string   `json:"producer"`
	ReleaseDate string   `json:"release_date"`
	RtScore     string   `json:"rt_score"`
	People      []string `json:"people"`
	Species     []string `json:"species"`
	Locations   []string `json:"locations"`
	Vehicles    []string `json:"vehicles"`
	URL         string   `json:"url"`
}

func getResponseData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	return contents
}
