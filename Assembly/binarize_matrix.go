package main

//BinarizeMatrix takes a matrix of values and a threshold.
//It binarizes the matrix according to the threshold.
//If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	numRows := len(mtx)

	b := make([][]int, numRows)
	for i := range b {
		b[i] = make([]int, numRows)
	}
	// lots of default zero values, which we don't need to set

	//range over matrix and set values
	for row := range b {
		for col := range b[row] {
			if mtx[row][col] >= threshold {
				//only fill in a value if it's bigger than mtx[col][row]
				if mtx[row][col] > mtx[col][row] {
					b[row][col] = 1
				} else if mtx[row][col] == mtx[col][row] {
					//always pick mtx[row][col] in a tie if row < col
					if row < col {
						b[row][col] = 1
					}
				}
			}
		}
	}

	return b
}
