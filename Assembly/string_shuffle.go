package main

import "math/rand"

//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	n := len(patterns)
	perm := rand.Perm(n)

	patterns2 := make([]string, n)

	for i := range patterns {
		index := perm[i]
		patterns2[i] = patterns[index]
	}

	return patterns2
}
