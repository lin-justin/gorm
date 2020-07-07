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
```go
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
