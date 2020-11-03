package Alignment

//GlobalScoreTable takes two strings, along with a match reward and mismatch and gap
//alignment penalties. It returns a 2-D array
//holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	scoringMatrix := make([][]float64, numRows)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]float64, numCols)
	}

	//first, penalize the 0-th row and column
	for j := range scoringMatrix {
		scoringMatrix[j][0] = float64(j) * (-gap)
	}
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = float64(j) * (-gap)
	}

	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//need an up, left, and a diag
			up := scoringMatrix[i-1][j] - gap //indel
			left := scoringMatrix[i][j-1] - gap
			diag := scoringMatrix[i-1][j-1]
			if str0[i-1] == str1[j-1] {
				//match!
				diag += match
			} else {
				//mismatch!
				diag -= mismatch
			}
			// now take max
			scoringMatrix[i][j] = MaxFloat(up, left, diag)
		}
	}

	return scoringMatrix

}

func MaxFloat(nums ...float64) float64 {
	if len(nums) == 0 {
		panic("Error: no values given!")
	}
	m := nums[0]
	for _, val := range nums {
		if val > m {
			m = val
		}
	}
	return m
}
