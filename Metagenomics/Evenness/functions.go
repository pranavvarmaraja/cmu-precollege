package Evenness

// Write a function SimpsonsIndex() with the following input and output
// Input: a frequency map sample
// Output: the Simpson’s index of sample. That is, the decimal sum of (n/N) squared
// over all species where n = # of individuals found for a given string/species
// and N is the total # of individuals.
func SimpsonsIndex(freq map[string]int) float64 {
	s := 0.0 // this makes s a float64 :)

	N := SumValues(freq)

	for _, val := range freq {
		n := val
		x := float64(n) / float64(N)
		x *= x // this gives probability of drawing pattern twice
		s += x
	}

	return s
}

//SumValues takes a frequency map of strings to ints as input.
//It sums all values in the frequency map, which it returns.
func SumValues(freq map[string]int) int {
	N := 0

	for _, val := range freq {
		N += val
	}

	return N
}
