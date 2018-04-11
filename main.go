package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

var left = flag.String("left", "", "left string")
var right = flag.String("right", "", "right string")

func main() {
	flag.Parse()

	if *left == "" || *right == "" {
		fmt.Println("we need left and right string to compare.")
		fmt.Println("usage: -left={string} -right={string}")
		return
	}

	fmt.Printf("Levenshtein distance is %d\n", Levenshtein(*left, *right))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// code from https://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Levenshtein_distance#Go
func Levenshtein(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)

	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0] // fj1 is the value of f[j - 1] in last iteration
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1) // delete & insert
			if cb != ca {
				mn = min(mn, fj1+1) // change
			} else {
				mn = min(mn, fj1) // matched
			}

			fj1, f[j] = f[j], mn // save f[j] to fj1(j is about to increase), update f[j] to mn
			j++
		}
	}

	return f[len(f)-1]
}
