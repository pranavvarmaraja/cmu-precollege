package main

//ConsensusFromMultAlign takes a collection columnCounts that holds, for each position of the multiple alignment, a map of DNA symbols to their number of occurrences. It also takes a voteThreshold.
//It returns a string corresponding to the consensus string of the multiple alignment.
func ConsensusFromMultAlign(columnCounts [](map[rune]int), voteThreshold float64) string {
	// columnCounts contains the multiple alignment data for a single contig.
	// Each col corresponds to a position in the alignment.
	// Each col contains [#A, #T, #C, #G] at that position.

	//we will represent the consensus string as an array of symbols to make conversion quick
	numColumns := len(columnCounts)
	consensusString := make([]rune, numColumns)

	//symbolCounts will hold total # of each symbol
	symbolCounts := make(map[rune]int)

	finalCol := 0

	total := 0
	maxNucleotideNum := 0 // e.g. number of As at current position
	maxNucleotide := 'X'
	vote := 0.
	// Iterate over all positions (cols) in the Multiple Alignment
	for col := range columnCounts {
		for _, sym := range []rune{'A', 'C', 'G', 'T'} {
			symbolCounts[sym] = columnCounts[col][sym]
			if symbolCounts[sym] > maxNucleotideNum {
				maxNucleotideNum = symbolCounts[sym]
				maxNucleotide = sym
			}
		}

		total = symbolCounts['A'] + symbolCounts['C'] + symbolCounts['G'] + symbolCounts['T']
		vote = float64(maxNucleotideNum) / float64(total)

		// columnCounts contains more columns than a single contig.
		// If we finish the contig, we start getting to the 0-columns of columnCounts.
		// And, we can stop.
		if total == 0 {
			finalCol = col
			break
		}

		if vote > voteThreshold {
			consensusString[col] = maxNucleotide
		} else {
			consensusString[col] = 'N'
		}

		// reset variables for the next position in the multiple alignment
		maxNucleotideNum = 0
		maxNucleotide = 'X'
	}

	return string(consensusString[:finalCol+1])
}
