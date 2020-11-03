package Clustal

import (
	"EvolutionaryTrees/Datatypes"
	"EvolutionaryTrees/Lib"
)

/*
 GenerateScoreTable : []string * []string -> [][]float
 REQUIRES : non-nil.
 ENSURES  : non-nil. for all i,j : float[i][j] >= 0.
                     float[len(align1)[0] + 1][len(align2)[0] + 1] is max.
 DESCRIP  : Populates path costs in manhattan graph between two
            alignments using a global, pairwise recurrence. Returns
            matrix representing these costs.
*/
func GenerateScoreTable(align1 Datatypes.Alignment, align2 Datatypes.Alignment,
	match float64, mismatch float64, gap float64, supergap float64) [][]float64 {

	numRows := len(align1[0]) + 1 //number of columns of first alignment
	numCols := len(align2[0]) + 1 //number of columns of second alignment

	//initialize scoring table
	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}

	// penalize the 0th row and column
	for j := 1; j < numCols; j++ {
		scoreTable[0][j] = float64(j) * (-supergap)
	}

	for i := 1; i < numRows; i++ {
		scoreTable[i][0] = float64(i) * (-supergap)
	}

	// traverse the interior of the scoring table.
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			// need to pick between the three incoming edges based on their weight + score of predecessor node.
			upValue := scoreTable[i-1][j] - supergap
			leftValue := scoreTable[i][j-1] - supergap
			diagValue := scoreTable[i-1][j-1] + SumPairsScore(align1, align2, i-1, j-1, match, mismatch, gap)

			// now pick the max of the three values for the current score table value.
			scoreTable[i][j] = Lib.MaxFloat(upValue, leftValue, diagValue)
		}
	}

	return scoreTable
}
