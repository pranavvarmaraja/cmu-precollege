package main

//OverlapScoringMatrix takes a collection of reads along with alignment penalties.
//It returns a matrix in which mtx[i][j] is the overlap alignment score of
//reads[i] with reads[j].
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	numReads := len(reads)

	mtx := InitializeSquareMatrix(numReads)

	//fill in the values of our matrix
	for row := range mtx {
		for col := range mtx[row] {
			if row != col { // leave the diagonals = 0
				mtx[row][col] = ScoreOverlapAlignment(reads[row], reads[col], match, mismatch, gap)
			}
		}
	}

	return mtx
}

//InitializeSquareMatrix takes an integer n and returns a 2-D slice of floats
//with default zero value.
func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}
