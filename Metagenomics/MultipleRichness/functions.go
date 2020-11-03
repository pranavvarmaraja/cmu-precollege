package MultipleRichness

import (
  "Metagenomics/Richness"
)

// Write the function RichnessMatrix() with the following input and output
// Input: an array of frequency maps samples
// Output: an slice whose i-th element is the richness of the i-th frequency map
func RichnessMatrix(allMaps []map[string]int) []int {
  numSamples := len(allMaps)
  a := make([]int, numSamples)

  // range over the maps, computing the richness of each and placing into a
  for i := range allMaps {
    a[i] = Richness.Richness(allMaps[i])
  }

  return a
}
