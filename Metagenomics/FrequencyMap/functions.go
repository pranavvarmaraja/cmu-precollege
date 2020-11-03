package FrequencyMap

// Write function FrequencyMap() with the following input and output
// Input: one collection of strings patterns
// Output: a frequency map of strings to their # of counts in patterns
func FrequencyMap(patterns []string) map[string]int {
	freqMap := make(map[string]int)

	for _, pattern := range patterns {
		freqMap[pattern]++
	}

	return freqMap
}
