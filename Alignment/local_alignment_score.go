package Alignment

//LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for local alignment with these penalties.
func LocalScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Error: String of zero length given to LocalAlignmentScoreTable.")
	}

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	// initialize scoring matrix -- would be great for subroutine
	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}

	// we want the 0-th row and column to all be zero.
	// good news: they already got this as default value.

	// now I am ready to range row by row over interior of table and apply the LA recurrence relation
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//apply the recurrence relation
			upValue := scoreTable[i-1][j] - gap   //indel
			leftValue := scoreTable[i][j-1] - gap //indel
			var diagonalWeight float64
			if str0[i-1] == str1[j-1] { //match!
				diagonalWeight = match
			} else { // mismatch!
				diagonalWeight = -mismatch
			}
			diagValue := scoreTable[i-1][j-1] + diagonalWeight
			scoreTable[i][j] = MaxFloat(upValue, leftValue, diagValue, 0.0)
		}
	}

	//only thing remaining is to allow for the free taxi rides to the end.
	//best local alignment will end at the maximum of all values in the score table.

	return scoreTable
}
