package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// API link to the Rick and Morty API
const API = "https://rickandmortyapi.com/api/character/"

func main() {
	contents := getResponseData(API)

	var characters AllCharacters
	json.Unmarshal(contents, &characters)

	var imgURL []string
	for _, character := range characters.Results {
		imgURL = append(imgURL, character.Image)
	}

	var wg sync.WaitGroup

	wg.Add(len(imgURL))

	for _, url := range imgURL {
		go func(url string) {
			defer wg.Done()
			tokens := strings.SplitAfter(url, "avatar/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)

			output, err := os.Create(filepath.Join("./images", fileName))
			if err != nil {
				log.Fatal("Error while creating", fileName, "-", err)
			}
			defer output.Close()

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer resp.Body.Close()
				_, err := io.Copy(output, resp.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("Done")
}

// Character type from the Rick and Morty API
type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	URL     string   `json:"url"`
	Created string   `json:"created"`
}

//AllCharacters struct to access all the character info in Rick and Morty
type AllCharacters struct {
	Results []Character `json:"results"`
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
