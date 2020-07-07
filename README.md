# GoGhibli

[Studio Ghibli API](https://ghibliapi.herokuapp.com/)

## Usage

```
$ go run main.go -film="My Neighbor Totoro"

Title: My Neighbor Totoro
------------------------
Description: Two sisters move to the country with their father in order to be closer to their hospitalized 
mother, and discover the surrounding trees are inhabited by Totoros, magical spirits of the forest. 
When the youngest runs away from home, the older sister seeks help from the spirits to find her.
------------------------
Director: Hayao Miyazaki
------------------------
Producer: Hayao Miyazaki
------------------------
Release Year: 1988
------------------------
Rotten Tomato Score: 93
------------------------
```

The default film is `Castle in the Sky`. The `-film` flag only accepts valid Studio Ghibli feature films. [Here](https://en.wikipedia.org/wiki/List_of_Studio_Ghibli_works#Feature_films) is a link to these films.

**The film must be enclosed in double quotation marks in order for the CLI to run, i.e. `-film="Feature Film of Interest"`.**

[Reference](https://github.com/mikicaivosevic/golang-json-client)

# GoRM

[Rick and Morty API](https://rickandmortyapi.com/)

This Go program is mainly for me to learn how to concurrently download multiple images from different URLs by connecting to a public API and parsing the nested JSON.

The [images folder](images) contains all the images downloaded from the Go program.

[Goroutine Reference](https://gist.github.com/nevermosby/b54d473ea9153bb75eebd14d8d816544)

[JSON Struct Reference](https://github.com/pitakill/rickandmortyapigowrapper/blob/master/character_structs.go)

# GoSM

A CLI program in Go that computes several string metric algorithms for two strings.

The algorithms used are Hamming Distance, Damerau-Levenshtein, Jaccard coefficient, and Jaccard distance.

## Usage
```
$ go run main.go -string1=kitten -string2=sitting -metric=hamming

The Hamming distance for kitten and sitting is: 2

$ go run main.go -string1=kitten -string2=sitting -metric=dl

The Damerau-Levenshtein distance for kitten and sitting is: 3
```
```
-metric=hamming // Hamming distance
-metric=dl // Damerau-Levenshtein
-metric=jc // Jaccard coefficient
-metric=jd // Jaccard distance
```
