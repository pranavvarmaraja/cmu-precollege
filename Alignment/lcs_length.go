package Alignment

//LCSLength takes two strings as input. It returns the length of a longest common
//subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}

	// let's use a subroutine to compute the matrix holding all the scores
	scoringMatrix := LCSScoreMatrix(str1, str2)

	numRows := len(scoringMatrix)
	numCols := len(scoringMatrix[numRows-1])

	return scoringMatrix[numRows-1][numCols-1]
}

//LCSScoreMatrix takes two strings as input. It returns a matrix holding
//length(i, j) for the count symbols recurrence.
func LCSScoreMatrix(str1, str2 string) [][]int {
	// what are array dimensions?

	numRows := len(str1) + 1
	numCols := len(str2) + 1

	//make our scoring matrix
	scoringMatrix := make([][]int, numRows)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]int, numCols)
	}
	//everybody is zero right now

	//range over values and apply the recurrence relation
	//row zero and column zero should all be zero anyway :)
	for row := 1; row < numRows; row++ {
		for col := 1; col < numCols; col++ {
			//apply recurrence relation
			up := scoringMatrix[row-1][col]
			left := scoringMatrix[row][col-1]
			diag := scoringMatrix[row-1][col-1]
			// when will we need to add 1 to diag??? When symbols match
			if str1[row-1] == str2[col-1] {
				diag++
			}
			scoringMatrix[row][col] = Max(up, left, diag)
		}
	}
	return scoringMatrix
}

//variadic function can take an arbitrary number of inputs

func Max(nums ...int) int {
	//the inputs to the function get created into an array called "nums"
	if len(nums) == 0 {
		panic("Error: no inputs given to Max().")
	}
	m := 0

	//range through nums, update if I find a bigger value
	for i, val := range nums {
		if i == 0 || val > m {
			m = val
		}
	}

	return m
}
