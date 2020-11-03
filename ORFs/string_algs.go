package main



// let's put our functions here.
func FrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)

	freqMap := FrequencyMap(text, k)

	// find the max value of frequency map
	m := MaxMap(freqMap)

	// what achieves the max?
	for pattern, val := range freqMap {
		if val == m {
			// frequent pattern found! append it to our list
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

func MaxMap(freqMap map[string]int) int {
	m := 0

	// range through map, and if something has higher value, update m!
	for pattern := range freqMap {
		if freqMap[pattern] > m {
			m = freqMap[pattern]
		}
	}
	// if all values in map were negative integers, this would return 0.
	// challenge: fix this bug so that it finds max value of any map of strings to ints.

	return m
}

//FrequencyMap takes a string text and an integer k. It produces a map
//of all k-mers in the string to their number of occurrences.
func FrequencyMap(text string, k int) map[string]int {
	// map declaration is analogous to slices
	// (we don't need to give an initial length)
	freq := make(map[string]int)
	n := len(text)
	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k]
		// if freqMap[pattern] doesn't exist, create it.  How do we do this?
		/*
		   // approach #1
		   _, exists := freq[pattern]
		   if exists == false {
		     // create this value
		     freqMap[pattern] = 1
		   } else {
		     // we already have a value in the map
		     freqMap[pattern]++
		   }
		*/
		// approach #2
		// this says, if freqMap[pattern] exists, add one to it
		// if freqMap[pattern] doesn't exist, create it with a default value of 0, and add 1.
		freq[pattern]++
	}
	return freq
}

func PatternCount(pattern, text string) int {
	hits := StartingIndices(pattern, text)
	return len(hits)
}

func StartingIndices(pattern, text string) []int {
	hits := make([]int, 0)

	// append every starting position of pattern that we find in text

	n := len(text)
	k := len(pattern)

	for i := 0; i < n-k+1; i++ {
		if text[i:i+k] == pattern {
			// hit found!
			hits = append(hits, i)
		}
	}

	return hits
}

func SkewArray(genome string) []int {
	n := len(genome)
	array := make([]int, n+1)

	for i := range genome {
		/*
		   array[i+1] = array[i] + something
		   something = -1, 0, 1 depending genome[i]
		*/
		if genome[i] == 'A' || genome[i] == 'T' {
			array[i+1] = array[i]
		} else if genome[i] == 'C' {
			array[i+1] = array[i] - 1
		} else if genome[i] == 'G' {
			array[i+1] = array[i] + 1
		}
	}

	return array
}

func ReverseComplement(text string) string {
	return Reverse(Complement(text))
}

//Reverse takes a string and returns the reversed string.
func Reverse(text string) string {
	n := len(text)
	symbols := make([]byte, n)
	for i := range text {
		symbols[i] = text[n-i-1]
	}
	return string(symbols)
}

func Complement(text string) string {
	// as with arrays, we can use "range"

	n := len(text)
	symbols := make([]byte, n)

	for i := range text {
		if text[i] == 'A' {
			symbols[i] = 'T'
		} else if text[i] == 'T' {
			symbols[i] = 'A'
		} else if text[i] == 'C' {
			symbols[i] = 'G'
		} else if text[i] == 'G' {
			symbols[i] = 'C'
		}
	}

	return string(symbols)
}
