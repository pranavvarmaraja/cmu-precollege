package Alignment

import "fmt"

type Alignment [2]string

//let's print out our work
func PrintAlignment(a Alignment) {
  fmt.Println(a[0])
  fmt.Println(a[1])
}

//GlobalAlignment takes two strings, along with match reward, and mismatch, and gap penalties.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str0, str1 string, match, mismatch, gap float64) Alignment {
  backtrack := GlobalBacktrack(str0, str1, match, mismatch, gap)
  //backtrack produces a matrix of backtracking pointers.
  //now infer the optimal alignment from this matrix
  return OutputGlobalAlignment(str0, str1, backtrack)
}

//GlobalBacktrack takes two strings and scoring parameters. It returns the backtrack
//matrix for global alignment.
func GlobalBacktrack(str0, str1 string, match, mismatch, gap float64) [][]string {
  if len(str0) == 0 || len(str1) == 0 {
    panic("Zero length strings given to GlobalBacktrack")
  }

  //initialize backtracking matrix
  numRows := len(str0) + 1
  numCols := len(str1) + 1

  backtrack := make([][]string, numRows)
  for i := range backtrack {
    backtrack[i] = make([]string, numCols)
  }

  //compute all the scores in the table
  scoreTable := GlobalScoreTable(str0, str1, match, mismatch, gap)

  //first, let's set backtracking pointers of 0-th row and column
  for j := 1; j < numCols; j++ {
    backtrack[0][j] = "LEFT"
  }
  for j := 1; j < numRows; j++ {
    backtrack[j][0] = "UP"
  }

  //traverse rest of scoring table, and figure out which previous value was used
  //to compute each score. This will let us set the backtracking pointer.
  for i := 1; i < numRows; i++ {
    for j := 1; j < numCols; j++ {
      if scoreTable[i][j] == scoreTable[i-1][j] - gap {
        backtrack[i][j] = "UP"
      } else if scoreTable[i][j] == scoreTable[i][j-1] - gap {
        backtrack[i][j] = "LEFT"
      } else {
        // better: check whether we have a match or mismatch
        // make sure that backtrack[i][j] = backtrack[i-1][j-1] + match if match
        // and = backtrack[i-1][j-1] - mismatch if mismatch
        // in any other case, panic
        backtrack[i][j] = "DIAG"
      }
    }
  }

  return backtrack
}

//OutputGlobalAlignment takes two strings and a backtracking matrix. It returns the optimal
//alignment of the two strings.
func OutputGlobalAlignment(str0, str1 string, backtrack [][]string) Alignment {
  var a Alignment
  //a[0] is top row, a[1] is bottom row
  //as I backtrack, I will chop symbols off the ends of the strings.

  for len(str0) > 0 || len(str1) > 0 {
    row := len(str0)
    col := len(str1)
    //these indices tell us where we are in backtrack matrix
    if backtrack[row][col] == "UP" {
      a[0] = string(str0[row-1]) + a[0]
      // what symbol do we put on the start of a[1]?
      a[1] = "-" + a[1]
      //chop off final symbol of str0
      str0 = str0[:len(str0)-1]
    } else if backtrack[row][col] == "LEFT" {
      a[0] = "-" + a[0]
      a[1] = string(str1[col-1]) + a[1]
      //chop off final symbol of str1
      str1 = str1[:len(str1)-1]
    } else {
      //lazy again :)
      a[0] = string(str0[row-1]) + a[0]
      a[1] = string(str1[col-1]) + a[1]
      // chop off symbol of both strings
      str0 = str0[:len(str0)-1]
      str1 = str1[:len(str1)-1]
    }
  }

  return a
}
