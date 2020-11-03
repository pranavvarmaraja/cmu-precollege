package Clustal

import (
	"EvolutionaryTrees/Datatypes"
)

/*
 ProgressiveAlign : Alignment * Alignment -> Alignment
 REQUIRES         : non-nil. strings in {A, T, C, G}*.
 ENSURES          : needleman alignment maximizing score.
 DESCRIP          : Performs a Smith-Waterman alignment between two existing
                    alignments using the ClustalW (1st. gen) scoring
                    heuristic.
*/
func ProgressiveAlign(align1 Datatypes.Alignment, align2 Datatypes.Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Datatypes.Alignment {

	scoreTable := GenerateScoreTable(align1, align2, match, mismatch, gap, supergap)

	// now generate the backtracking matrix
	backtrack := GenerateBacktrackMatrix(align1, align2, scoreTable, match, mismatch, gap, supergap)

	return OutputAlignment(align1, align2, backtrack)

}

func OutputAlignment(align1 Datatypes.Alignment, align2 Datatypes.Alignment, backtrack [][]string) Datatypes.Alignment {
	numStr1 := len(align1)
	numStr2 := len(align2)

	//first, initialize alignment

	newAlignment := make(Datatypes.Alignment, numStr1+numStr2)

	//now, add in all strings into alignment
	for i, str := range align1 {
		newAlignment[i] = Backtracker(str, backtrack, "side")
	}

	for j, str := range align2 {
		newAlignment[j+numStr1] = Backtracker(str, backtrack, "top")
	}

	return newAlignment
}

func Backtracker(str string, backtrack [][]string, orientation string) string {
	// we are producing a single row of the combined alignment.

	aligned := ""

	//where are we in the matrix? at the end ...
	row := len(backtrack) - 1
	col := len(backtrack[0]) - 1

	//continue backtracking until the source

	for row != 0 || col != 0 {
		k := len(str)
		if backtrack[row][col] == "UP" {
			//taking only something from alignment 1
			// am I in alignment 1 or 2?
			if orientation == "side" {
				aligned = string(str[k-1]) + aligned
				// shorten current string by a single symbol
				str = str[:k-1]
			} else if orientation == "top" {
				//gap symbol
				aligned = "-" + aligned
			}
			// decrease row index
			row--
		} else if backtrack[row][col] == "LEFT" {
			//only alignment 2
			if orientation == "side" {
				//gap symbol
				aligned = "-" + aligned
			} else if orientation == "top" {
				//take a symbol from current string
				aligned = string(str[k-1]) + aligned
				// shorten current string by a single symbol
				str = str[:k-1]
			}
			// decrease col by 1
			col--
		} else {
			//case of both alignments
			// always take symbol, regardless of being in one alignment or the other
			aligned = string(str[k-1]) + aligned
			str = str[:k-1]
			row--
			col--
		}
	}

	return aligned
}

func GenerateBacktrackMatrix(align1, align2 Datatypes.Alignment, scoreTable [][]float64, match, mismatch, gap, supergap float64) [][]string {

	//initialize our 2-D slice of strings to store all the pointers

	numRows := len(align1[0]) + 1
	numCols := len(align2[0]) + 1

	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	//set the backtracking pointers based on scores

	//0-th row and column are EZ
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}

	for i := 0; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	backtrack[0][0] = "START" // not necessary but gives us a starting point

	// range over inside of table and set based on which edge was used
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			if scoreTable[i][j] == scoreTable[i][j-1]-supergap {
				backtrack[i][j] = "LEFT"
			} else if scoreTable[i][j] == scoreTable[i-1][j]-supergap {
				backtrack[i][j] = "UP"
			} else {
				//hopefully scoreTable[i][j] = sum of pairs score + scoretable[i-1][j-1]
				backtrack[i][j] = "DIAG"
			}
		}
	}

	return backtrack

}
