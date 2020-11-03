package main

//MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
//It returns adjacency lists of reads; edges are only included
//in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	adjacencyList := make(map[string][]string)

	// fill in details of the function here.
	// first, construct our overlap scoring matrix
	mtx := OverlapScoringMatrix(reads, match, mismatch, gap)

	//next, binarize
	b := BinarizeMatrix(mtx, threshold)

	//now, traverse matrix and fill in values of the adjacency list
	for row := range b {
		for col := range b[row] {
			//only take an action if current matrix value is 1
			if b[row][col] == 1 {
				adjacencyList[reads[row]] = append(adjacencyList[reads[row]], reads[col])
			}
		}
	}

	return adjacencyList
}
