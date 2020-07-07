// A CLI program that computes various string metric algorithms
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	firstWordPtr := flag.String("string1", "karolin", "a string")
	secondWordPtr := flag.String("string2", "kathrin", "a string")

	stringMetricAlgorithm := flag.String("metric", "hamming", "a string metric algorithm\noptions are:\nhamming, dl, jc, jd")

	flag.Parse()

	switch *stringMetricAlgorithm {
	case "dl":
		result := DamerauLevenshteinDistance(*firstWordPtr, *secondWordPtr)
		fmt.Printf("The Damerau-Levenshtein distance for %s and %s is: %v\n", *firstWordPtr, *secondWordPtr, result)
	case "jc":
		result := JaccardCoefficient(*firstWordPtr, *secondWordPtr)
		fmt.Printf("The Jaccard Coefficient for %s and %s is: %.4f\n", *firstWordPtr, *secondWordPtr, result)
	case "jd":
		result := JaccardDistance(*firstWordPtr, *secondWordPtr)
		fmt.Printf("The Jaccard distance for %s and %s is: %.4f\n", *firstWordPtr, *secondWordPtr, result)
	default:
		result := HammingDistance(*firstWordPtr, *secondWordPtr)
		fmt.Printf("The Hamming distance for %s and %s is: %v\n", *firstWordPtr, *secondWordPtr, result)
	}
}

// HammingDistance calculates the minimum number of substitutions required to
// change one string into the other, or the minimum number of errors that could
// have transformed one string into the other.
// Reference: https://en.wikipedia.org/wiki/Hamming_distance
func HammingDistance(s1, s2 string) int {
	distCounter := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distCounter++
		}
	}
	return distCounter
}

// getMin gets the minimum value from a slice of ints
func getMin(vals []int) int {
	m := 0
	for i, v := range vals {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}

// generateRange makes a slice of ints from 0 to the limit
func generateRange(limit int) []int {
	var r []int
	for i := 0; i < limit; i++ {
		r = append(r, i)
	}
	return r
}

// DamerauLevenshteinDistance calculates the optimal string alignment
// using the Damerau-Levenshtein formula
// Reference: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance#Optimal_string_alignment_distance
//			  https://github.com/dansackett/levenshtein/blob/master/levenshtein.go
func DamerauLevenshteinDistance(source, target string) int {
	sourceRange := generateRange(len(source) + 1)
	targetRange := generateRange(len(target) + 1)
	matrix := make([][]int, len(source)+1)

	// Create initial matrix
	for _, s := range sourceRange {
		matrix[s] = make([]int, len(target)+1)
		for _, t := range targetRange {
			if t == 0 {
				matrix[s][t] = s
			} else {
				matrix[s][t] = t
			}
		}
	}

	// Determine cost for insertion, deletion, subsitution, and
	// transposition and use the minimum value as the true editing
	// distance for the given character position in the matrix
	for _, s := range sourceRange[1:] {
		for _, t := range targetRange[1:] {
			subTransCost := 0

			if source[s-1] == target[t-1] {
				subTransCost = 0
			} else {
				subTransCost = 1
			}

			deleteDistance := matrix[s-1][t] + 1
			insertDistance := matrix[s][t-1] + 1
			subsitutionDistance := matrix[s-1][t-1] + subTransCost

			matrix[s][t] = getMin([]int{deleteDistance, insertDistance, subsitutionDistance})

			if s > 1 && t > 1 && source[s-1] == target[t-2] && source[s-2] == target[t-1] {
				transpositionDistance := matrix[s-2][t-2] + subTransCost
				matrix[s][t] = getMin([]int{matrix[s][t], transpositionDistance})
			}
		}
	}
	return matrix[len(source)][len(target)]
}

// Reference: https://github.com/hueyjj/jaccard-string
func intersect(s1, s2 string) string {
	var intersection strings.Builder
	set := make(map[rune]bool)

	for _, char := range s1 {
		if _, ok := set[char]; !ok {
			set[char] = true
		}
	}

	for _, char := range s2 {
		if _, ok := set[char]; ok && set[char] {
			intersection.WriteRune(char)
			set[char] = false
		}
	}
	return intersection.String()
}

func union(s1, s2 string) string {
	var intersection strings.Builder
	set := make(map[rune]bool)

	for _, char := range s1 {
		if _, ok := set[char]; !ok {
			set[char] = true
			intersection.WriteRune(char)
		}
	}

	for _, char := range s2 {
		if _, ok := set[char]; !ok {
			set[char] = true
			intersection.WriteRune(char)
		}
	}
	return intersection.String()
}

// JaccardCoefficient measures similarity between finite sample sets and is defined
// as the size of the intersection divided by the size of the unionof the sample sets
func JaccardCoefficient(s1, s2 string) float64 {
	return float64(len(intersect(s1, s2))) / float64(len(union(s1, s2)))
}

// JaccardDistance measures dissimlarity between sample set and is complementary to the
// Jaccard coefficient and is obtained by subtracting the Jaccard coefficient from 1.
func JaccardDistance(s1, s2 string) float64 {
	jaccardCoefficient := JaccardCoefficient(s1, s2)
	return 1 - jaccardCoefficient
}
