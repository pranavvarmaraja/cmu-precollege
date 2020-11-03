package Alignment

//LongestCommonSubsequence takes two strings str0 and str1 as input.
//It should return a longest common subsequence of the two strings.
func LongestCommonSubsequence(str0, str1 string) string {
	// we need an array that will store backtracking pointers
	backtrack := LCSBacktrack(str0, str1)
	return OutputLCS(str0, str1, backtrack)
}

//OutputLCS takes two strings and an array of backtracking pointers.
//It uses these to produce the longest common subsequence of the strings.
func OutputLCS(str0, str1 string, backtrack [][]string) string {
	s := ""

	//I'm going to chop off symbols as we go. When I run out of symbols in one string, I'm done
	for len(str0) > 0 || len(str1) > 0 {
		row := len(str0)
		col := len(str1)
		if backtrack[row][col] == "UP" {
			//chop off one symbol of str0
			str0 = str0[:len(str0)-1]
		} else if backtrack[row][col] == "LEFT" {
			str1 = str1[:len(str1)-1]
		} else if backtrack[row][col] == "DIAG" {
			//we could be at a match, or at a mismatch
			if str0[len(str0)-1] == str1[len(str1)-1] {
				//match
				//add a symbol to the start of s
				s = string(str0[len(str0)-1]) + s
			}
			//now chop off a symbol of both str0 and str1
			str0 = str0[:len(str0)-1]
			str1 = str1[:len(str1)-1]
		} else {
			panic("Error: invalid value in backtrack array.")
		}
	}

	return s
}

//LCSBacktrack takes two strings and produces a matrix of backtracking "pointers"
//for every node in an LCS table.
func LCSBacktrack(str0, str1 string) [][]string {
	//we want a backtracking "pointer" for every node.
	numRows := len(str0) + 1
	numCols := len(str1) + 1

	//make the matrix
	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	//generate the scoring table using the function we already wrote
	scoringMatrix := LCSScoreMatrix(str0, str1)

	//fill the matrix

	//first, the easiest backtracking pointers to set are the 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	//traverse the rest of the table and set pointers based on neighboring values
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//check the value that was used to give me scoringMatrix[i][j]
			if scoringMatrix[i][j] == scoringMatrix[i-1][j] {
				backtrack[i][j] = "UP"
			} else if scoringMatrix[i][j] == scoringMatrix[i][j-1] {
				backtrack[i][j] = "LEFT"
			} else { // assume that the scoring matrix is correct
				backtrack[i][j] = "DIAG"
			}
		}
	}

	return backtrack
}
