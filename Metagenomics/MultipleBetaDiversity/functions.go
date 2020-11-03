package MultipleBetaDiversity

import (
	"Jaccard"
	"Metagenomics/BrayCurtis"
)

// Write the functions BetaDiversityMatrix() with the following input and output
// Input: a slice of frequency maps samples and a distance metric ("Bray-Curtis" or "Jaccard")
// Output: a distance matrix where D_i,j is the distance between
// sample_i and sample_j according to the given distance metric
func BetaDiversityMatrix(allMaps []map[string]int, distMetric string) [][]float64 {
	numSamples := len(allMaps)
	mtx := InitializeSquareMatrix(numSamples)

	//fill in values of mtx
	// range over rows
	for i := 0; i < numSamples; i++ {
		// range over columns
		for j := i + 1; j < numSamples; j++ { // i is never j
			//set my values
			//two cases depending on distMetric
			if distMetric == "Bray-Curtis" {
				d := BrayCurtis.BrayCurtisDistance(allMaps[i], allMaps[j])
				mtx[i][j] = d
				//because of symmetry, this is mtx[j][i] as well
				mtx[j][i] = d
			} else if distMetric == "Jaccard" {
				d := Jaccard.JaccardDistance(allMaps[i], allMaps[j])
				mtx[i][j] = d
				//because of symmetry, this is mtx[j][i] as well
				mtx[j][i] = d
			} else {
				panic("Error: invalid distance metric given to MultipleBetaDiversity")
			}
		}
	}

	return mtx
}

//InitializeSquareMatrix takes an integer n and returns n x n
//decimal matrix of zeroes
func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n) // n rows
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}

	return mtx
}
