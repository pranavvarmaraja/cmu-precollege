package MultipleEvenness

import (
  "Metagenomics/Evenness"
)

// Write the function SimpsonsMatrix() with the following input and output
// Input: an array of frequency maps samples
// Output: an slice  i-th element is the Simpson's index of the i-th frequency map
func SimpsonsMatrix(allMaps []map[string]int) []float64 {
  a := make([]float64, len(allMaps))

  for i := range allMaps {
    a[i] = Evenness.SimpsonsIndex(allMaps[i])
  }

  return a
}
