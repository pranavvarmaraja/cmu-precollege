package Alignment

//LocalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score local alignment of the strings corresponding to these penalties.
func LocalAlignment(str0, str1 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
  backtrack := LocalBacktrack(str0, str1, match, mismatch, gap)
	// backtrack will compute scores using recurrences and give me all the backtracking pointers as a 2-D array.
	optAlignment, start0, end0, start1, end1 := OutputLocalAlignment(str0, str1, backtrack)
	// this function will use the backtracking info to construct our alignment
	return optAlignment, start0, end0, start1, end1
}

func OutputLocalAlignment(str0, str1 string, backtrack [][]([2]int)) (Alignment, int, int, int, int) {
  var a Alignment

  numRows := len(backtrack)
  numCols := len(backtrack[0])

  endingPoint := backtrack[numRows-1][numCols-1]
  //where should I jump?
  endRow := endingPoint[0]
  endCol := endingPoint[1]

  //now backtrack from my position
  currRow := endRow
  currCol := endCol

  startRow := 0
  startCol := 0

  for currRow != 0 || currCol != 0 {
    //check possible cases for what the backtrack could be
    if backtrack[currRow][currCol] == [2]int{currRow-1, currCol} {
      //backtrack up one row
      a[0] = string(str0[currRow-1]) + a[0]
      a[1] = "-" + a[1]
      //now lower the current row
      currRow--
    } else if backtrack[currRow][currCol] == [2]int{currRow, currCol-1} {
      //backtrack over one column
      a[0] = "-" + a[0]
      a[1] = string(str1[currCol-1]) + a[1]
      //now lower the current column
      currCol--
    } else if backtrack[currRow][currCol] == [2]int{currRow-1, currCol-1} {
      a[0] = string(str0[currRow-1]) + a[0]
      a[1] = string(str1[currCol-1]) + a[1]
      //now lower the current row and col
      currRow--
      currCol--
    } else if backtrack[currRow][currCol] == [2]int{0,0} {
      //backtrack to source
      //we are done :)
      startRow = currRow
      startCol = currCol
      //lower currRow, currCol to zero
      currRow = 0
      currCol = 0
    } else {
      panic("Oh no")
    }
  }

  return a, startRow, endRow, startCol, endCol
}

func LocalBacktrack(str0, str1 string, match, mismatch, gap float64) [][]([2]int) {
  numRows := len(str0) + 1
  numCols := len(str1) + 1

  backtrack := make([][]([2]int), numRows)
  for i := range backtrack {
    backtrack[i] = make([]([2]int), numCols)
  }

  // let's grab the scores from our scoring table
  scoreTable := LocalScoreTable(str0, str1, match, mismatch, gap)

  //next, let's point the pointers in 0-th row and column at (0,0)
  for j := 1; j < numCols; j++ {
    backtrack[0][j] = [2]int{0, 0}
  }
  for i := 1; i < numRows; i++ {
    backtrack[i][0] = [2]int{0, 0}
  }

  //traverse interior of the table
  for i := 1; i < numRows; i++ {
    for j := 1; j < numCols; j++ {
      //check which value was used
      if scoreTable[i][j] == 0 {
        backtrack[i][j] = [2]int{0, 0} // go back to source
      } else if scoreTable[i][j] == scoreTable[i-1][j]-gap {
				// indel
				backtrack[i][j] = [2]int{i - 1, j}
			} else if scoreTable[i][j] == scoreTable[i][j-1]-gap {
				backtrack[i][j] = [2]int{i, j - 1}
			} else { // now we know that we backtrack along a diagonal
				backtrack[i][j] = [2]int{i - 1, j - 1}
			}
    }
  }

  // at the very end, determine where the max value in the table is, and
  //update the backtracking pointer of the sink to point at this value
  endRow, endCol := FindMaxIndices(scoreTable)
  backtrack[numRows-1][numCols-1] = [2]int{endRow, endCol}

  return backtrack

}

//FindMaxIndices takes a matrix and returns the indices (row, col) where the maximum value is found.
func FindMaxIndices(mtx [][]float64) (int, int) {
	bestRow, bestCol := 0, 0
	m := mtx[bestRow][bestCol]

	// if I can beat the value at bestRow, bestCol, then I update
	for row := range mtx {
		for col := range mtx[row] {
			if mtx[row][col] > m { // better current max found!
				// update m, along with the best indices
				m = mtx[row][col]
				bestRow = row
				bestCol = col
			}
		}
	}
  return bestRow, bestCol
}
