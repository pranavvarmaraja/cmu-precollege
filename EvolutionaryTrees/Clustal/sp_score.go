package Clustal

import (
	"EvolutionaryTrees/Datatypes"
)

/*
 SumPairsScore : []string * []string * int * int -> int
 REQUIRES      : non-nil. idx1 < len(align1). idx2 < len(align2).
 ENSURES       : result >= 0; result = n, where n is the number of pairwise
                                          similarities btw alignments.
 DESCRIP       : Calculates the number of pairwise similar DNA characters between
                 at a particular index of each alignment.
                 ex. [A, T, T] [T, A] = 3
*/
func SumPairsScore(align1 Datatypes.Alignment, align2 Datatypes.Alignment,
	idx1 int, idx2 int, match float64, mismatch float64, gap float64) float64 {

	var totalScore = 0.0
	for _, str1 := range align1 {
		for _, str2 := range align2 {

			var char1 = str1[idx1]
			var char2 = str2[idx2]

			if char1 == '-' && char2 == '-' {
				totalScore += 0
			} else if char1 == '-' || char2 == '-' {
				totalScore -= gap
			} else if char1 == char2 {
				totalScore += match
			} else {
				totalScore -= mismatch
			}

		}
	}
	return totalScore
}
