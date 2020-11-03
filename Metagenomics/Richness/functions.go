package Richness

// Write function Richness() with the following input and output
// Input: a frequency map sample
// Output: the richness of sample. That is, the number of strings in the
// frequency map corresponding to sample

//Richness takes a frequency map. It returns the richness of the frequency map
//(i.e., the number of keys in the map corresponding to nonzero values.)
func Richness(sample map[string]int) int {
	c := 0 // holds the count of number of species seen

	for _, val := range sample {
		if val < 0 {
			panic("Error: negative value in frequency map given to Richness()")
		}
		if val > 0 {
			c++
		}
	}

	return c
}
