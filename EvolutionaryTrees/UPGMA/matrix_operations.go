package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
)

/*
  FindMinElement : DistanceMatrix * ->  int * int * float
  REQUIRES       : len(mtx) == len(mtx[len(mtx) - 1]) ; mtx[i][j] == m[j][i]
              		 mtx[i][i] == 0 ; mtx[i][j] != 0
  ENSURES        : mtx[row][col] == mtx[col][row] is the min elt in mtx
                   col > row (looking at top half of mtx)
  DESCRIP        : Returns the row index, column index, and value corresponding
                   to a minimum element.
*/
func FindMinElement(mtx Datatypes.DistanceMatrix) (int, int, float64) {
	if len(mtx) <= 1 || len(mtx[0]) <= 1 {
		panic("One row or one column!")
	}

	// can now assume that matrix is at least 2 x 2
	row := 0
	col := 1
	minVal := mtx[row][col]

	// range over matrix, and see if we can do better than minVal.
	for i := 0; i < len(mtx)-1; i++ {
		// start column ranging at i + 1
		for j := i + 1; j < len(mtx[i]); j++ {
			// do we have a winner?
			if mtx[i][j] < minVal {
				// update all three variables
				minVal = mtx[i][j]
				row = i
				col = j
				// col will still always be > row.
			}
		}
	}
	return row, col, minVal
}

/*
  DelRowCol : DistanceMatrix * int * int -> DistanceMatrix
  REQUIRES  : len(mtx) == len(mtx[len(mtx) - 1]) ; mtx[i][j] == m[j][i]
              row < len(mtx) ; col < len(mtx[0]) ; mtx[i][i] == 0 ;
							mtx[i][j] != 0
  ENSURES   : len(\result) == len(mtx)-1 ; len(\result[0]) == len(mtx[0]) - 1
  DESCRIP   : Given a distance matrix and two indices, deletes the rows and
              cols at each respective index and returns the resulting matrix.
*/
func DelRowCol(mtx Datatypes.DistanceMatrix, row, col int) Datatypes.DistanceMatrix {
	// first, let's delete appropriate rows
	// remember that col > row, we should delete col-th row first
	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)

	//now, delete columns row and col as well.
	for i := range mtx {
		mtx[i] = append(mtx[i][:col], mtx[i][col+1:]...)
		mtx[i] = append(mtx[i][:row], mtx[i][row+1:]...)
	}

	return mtx
}

/*
  AddRowCol   : DistanceMatrix * []*Node * int * int -> DistanceMatrix
  REQUIRES    : len(mtx) == len(mtx[len(mtx) - 1]) ; mtx[i][j] == m[j][i]
                row < len(mtx) ; col < len(mtx) ; len(clusters) == len(mtx)
  ENSURES     : len(\result) == len(mtx) + 1 ; len(\result[0]) == len(mtx[0]) + 1
  DESCRIP     : Given a DistanceMatrix, a slice of current clusters,
                and a row/col index (col > row),
                Returns the matrix corresponding to "gluing" clusters[row]
                and clusters[col] together and forming a new row/col of the
                matrix (no deletions yet).
*/
func AddRowCol(mtx Datatypes.DistanceMatrix,
	clusters []*Datatypes.Node, row, col int) Datatypes.DistanceMatrix {

	n := len(mtx)
	newRow := make([]float64, n+1) // last element will be 0 by default :)

	// all values 0.0 by default, let's set the values that need to be set.
	for j := 0; j < n; j++ {
		if j != row && j != col {
			// now compute newRow[j]
			// weighted average based on number of elements in each cluster
			size1 := CountLeaves(clusters[row])
			size2 := CountLeaves(clusters[col])

			newRow[j] = (float64(size1)*mtx[row][j] + float64(size2)*mtx[col][j]) / float64(size1+size2)
		}
	}

	//let's append new row to matrix
	mtx = append(mtx, newRow)

	//we need to add last column as well to the matrix
	for i := 0; i < n; i++ {
		mtx[i] = append(mtx[i], newRow[i])
	}

	return mtx
}
