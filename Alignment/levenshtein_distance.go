package Alignment

//EditDistance takes two strings as input. It returns the Levenshtein distance
//between the two strings; that is, the minimum number of substitutions, insertions, and deletions
//needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}

	scoringMatrix := EditDistanceMatrix(str1, str2)

	numRows := len(scoringMatrix)
	numCols := len(scoringMatrix[numRows-1])
	return scoringMatrix[numRows-1][numCols-1]
}

//EditDistanceMatrix takes two strings as input. It returns a matrix
//for which M(i, j) is the edit distance between the first i symbols of one string
//and the first j symbols of the other.
func EditDistanceMatrix(str1, str2 string) [][]int {
	//perhaps panic if either string has zero length

	numRows := len(str1) + 1
	numCols := len(str2) + 1

	scoringMatrix := make([][]int, numRows)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]int, numCols)
	}

	//right now, everyone in scoring matrix is zero.

	// set 0-th row and column first
	for j := range scoringMatrix {
		scoringMatrix[j][0] = j // sets 0-th column
	}
	//now set row 0
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = j
	}

	//only thing that really differs from finding LCS length is the recurrence itself.
	for row := 1; row < numRows; row++ {
		for col := 1; col < numCols; col++ {
			//current value should be minimum of three things
			up := scoringMatrix[row-1][col] + 1
			left := scoringMatrix[row][col-1] + 1
			diag := scoringMatrix[row-1][col-1]
			// add 1 to diag if we are at a mismatch
			if str1[row-1] != str2[col-1] {
				diag++
			}

			scoringMatrix[row][col] = Min(up, left, diag)
		}
	}

	return scoringMatrix
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("Error: no values given to Min()!")
	}
	m := nums[0]

	for _, val := range nums {
		if val < m {
			m = val
		}
	}

	return m
}
