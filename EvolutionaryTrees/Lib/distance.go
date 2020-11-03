package Lib

import (
	"fmt"
	"EvolutionaryTrees/Datatypes"
)

/************************************************
  CORE LOGIC
************************************************/

/* CalcDistancematrix : []string -> DistanceMatrix
   REQUIRES           : true
	 ENSURES            : len(\result) == len(\result[0]) ; \result[i][j] == \result[j][i]
												len(\result) == len(dnaStrings)
	 DESCRIP            : Given a string slice, creates a pairwise distance matrix
	                      between all strings. Uses alignment distance heuristic, better
												known as Levenshtein.
 */
func CalcDistanceMatrix(dnaStrings []string) Datatypes.DistanceMatrix {
	numSamples := len(dnaStrings)
	mtx := InitializeMatrix(numSamples, numSamples)

	// Progress Indicator
	n := float64(((numSamples*numSamples) + numSamples) / 2)
	k := 1.0

	for i := 0; i < numSamples; i++ {
		for j := i; j < numSamples; j++ {
			d := LevDistance(dnaStrings[i], dnaStrings[j])
			mtx[i][j] = d
			mtx[j][i] = d

			// Progress Indicator
			fmt.Println(fmt.Sprintf("%.2f", float64((k/n) * 100)) + "%")
			k = k+1
		}
	}
	return mtx
}

/* LevDistance : []string -> DistanceMatrix
   REQUIRES    : true
	 ENSURES     : len(\result) == len(\result[0]) ; \result[i][j] == \result[j][i]
								 len(\result) == len(dnaStrings)
	 DESCRIP     : Given a string slice, creates a pairwise distance matrix
	               between all strings. Uses alignment distance heuristic, better
								 known as Levenshtein.
 */
func LevDistance(seq1 string, seq2 string) float64 {
	var seq1Size = len(seq1) + 1
	var seq2Size = len(seq2) + 1
	var matrix = InitializeMatrix(seq1Size, seq2Size)

	for i := 0 ; i < seq1Size ; i++ {
		matrix[i][0] = float64(i)
	}
	for j := 0 ; j < seq2Size ; j++ {
		matrix[0][j] = float64(j)
	}

	for i := 1 ; i < seq1Size ; i++ {
		for j := 1 ; j < seq2Size ; j++ {
			if seq1[i-1] == seq2[j-1] {
				matrix[i][j] =
					MinFloat(
					  matrix[i-1][j] + 5,
					  matrix[i-1][j-1],
					  matrix[i][j-1] + 5)
			} else {
				matrix[i][j] =
					MinFloat(
						matrix[i-1][j] + 5,
						matrix[i-1][j-1] + 1,
						matrix[i][j-1] + 5)
			}
		}
	}
  return matrix[seq1Size-1][seq2Size-1]

}

/************************************************
  MISCELLANEOUS
************************************************/

func MinFloat(nums ...float64) float64 {
	m := 0.0
	// nums gets converted to an array
	for i, val := range nums {
		if val < m || i == 0 {
			// update m
			m = val
		}
	}
	return m
}

func InitializeMatrix(m int, n int) Datatypes.DistanceMatrix {
	mtx := make([][]float64, m)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}
