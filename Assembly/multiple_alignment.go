package main

//MultAlignmentFromContig takes a contig as a collection of strings coming from a non-branching path.
//It returns a multiple alignment, as an array of columns. Each column is a
//map from symbols to their counts in the column.
//It applies a greedy algorithm in which we attempt all possible no-gap alignments of
//the next string in the path against the current multiple alignment.
func MultAlignmentFromContig(contig []string) [](map[rune]int) {
	// index   -> node position in contig, e.g. index=3 means the 3rd node in path
	// value   -> start position of the node in the multiple alignment

	match := 1
	mismatch := 1

	//startingPositions is a slice holding the starting position
	//in the multiple alignment of each string
	startingPositions := make([]int, len(contig))

	//numChars will sum the number of characters in the path
	numChars := 0
	for i := 0; i < len(contig); i++ {
		numChars = numChars + len(contig[i])
	}

	// mutliple alighment contents -columnCounts- is a slice of maps
	// e.g. columnCounts[0]['A'] tells how many A's are in position 0
	//we initialize this array to have numChars values (probably too many)
	columnCounts := make([](map[rune]int), numChars)

	// now make all the internal maps
	for i := range columnCounts {
		columnCounts[i] = make(map[rune]int)

		//this is redundant but helpful to see the counts get assigned
		columnCounts[i]['A'] = 0
		columnCounts[i]['C'] = 0
		columnCounts[i]['G'] = 0
		columnCounts[i]['T'] = 0
	}

	//str0 is first string in the path
	str0 := contig[0]

	// Fill into columnCounts the nucelotide content for str0
	for i, sym := range str0 {
		columnCounts[i][sym]++
	}

	// Iterate over all remaining strings in this contig.
	// Provide a starting position for each (in the startingPositions slice)
	for i := 1; i < len(contig); i++ {
		predecessorLength := len(contig[i-1])
		currentString := contig[i]

		columnCounts, startingPositions = UpdateAlignment(columnCounts, startingPositions, currentString, i-1, predecessorLength, match, mismatch)
	}

	finalLength := len(contig[len(contig)-1])

	//trim multiple alignment to rid ourselves of zeroes
	return columnCounts[:startingPositions[len(contig)-1]+finalLength]
}

//UpdateAlignment takes a column count as a slice of character count maps, along with a few other parameters
//including a slice containing the starting position of each contig.
//It returns the updated column count and the starting positions after threading the current string given into the
//multiple alignment.
func UpdateAlignment(columnCounts [](map[rune]int), startingPositions []int, currentString string, j, predecessorLength, match, mismatch int) ([](map[rune]int), []int) {
	// j is the index of string in the contig most recently added to the multiple alignment
	// The slice columnCounts holds the nucleotide amounts for every position in the multiple
	// alignment so far.

	//will need alphabet
	alphabet := []rune{'A', 'C', 'G', 'T'}

	// startingPos is the starting position of the previous string in the multiple alignment
	startingPos := startingPositions[j]

	// For every starting position of currentString, compute a potential new columnCounts.
	// We only need to consider the rows of columnCounts starting at startingPos when considering
	// where to place currentString.
	// Similarly, we only need to consider up through startingPos + len(str0) + len(currentString)
	newColumnCounts := make([][](map[rune]int), predecessorLength)

	for i := range newColumnCounts {

		newColumnCounts[i] = make([]map[rune]int, predecessorLength+len(currentString))
		for j := range newColumnCounts[i] {
			rowMap := make(map[rune]int)
			//copy values of startinPosition+f into rowSlice
			for symbol := range columnCounts[startingPos+j] {
				rowMap[symbol] = columnCounts[startingPos+j][symbol]
			}

			newColumnCounts[i][j] = rowMap
		}
	}

	// scoreList is a list of scores, one for each offset of currentString relative to str0
	// For each nucleotide in currentString, compute match score with all nucleotides at
	// the same position (above it) in the multiple alignment, i.e. in newColumnCounts
	// score only considers match/mismatch. if nothing is above, no indel penalty.
	// Sum these scores over all nucleotides in currentString.
	scoreList := make([]int, predecessorLength)

	//range over length of predecessor string to update scoreList at each possible position
	for start := 0; start < predecessorLength; start++ {

		symbolScores := make(map[rune]int)

		symbolScores['A'] = 0
		symbolScores['C'] = 0
		symbolScores['G'] = 0
		symbolScores['T'] = 0

		// update newColumnCounts for this starting position
		// For each nucleotide in currentString, increment appropriate entry of newColumnCounts
		for i, currentSymbol := range currentString {
			newColumnCounts[start][start+i][currentSymbol]++

			//range over alphabet and assign match/mismatch penalties to newColumnCounts
			for _, symbol := range alphabet {
				if symbol == currentSymbol { //match
					symbolScores[symbol] += match*newColumnCounts[start][start+i][symbol] - 1
				} else { //mismatch
					symbolScores[currentSymbol] -= mismatch * newColumnCounts[start][start+i][symbol]
				}
			}
		}

		//now set scoreList equal to the sum of the values in symbolScores
		scoreList[start] = symbolScores['A'] + symbolScores['C'] + symbolScores['G'] + symbolScores['T']
	}

	// choose the best scoring start position
	var max int
	bestStart := 0
	for index, value := range scoreList {
		if index == 0 || value > max {
			bestStart = index
			max = value
		}
	}

	// UPDATE startingPositions for new node
	// starting position relative to position 0 in the multiple alignment, i.e. node 1
	shiftedStart := bestStart + startingPos

	startingPositions[j+1] = shiftedStart

	// UPDATE columnCounts- write newColumnCounts[bestStart] into columnCounts
	for w := startingPos; w < startingPos+predecessorLength+len(currentString); w++ {

		for _, symbol := range alphabet {
			columnCounts[w][symbol] = newColumnCounts[bestStart][w-startingPos][symbol] //newColumnCounts[bestStart][w-p][0]
		}
	}

	return columnCounts, startingPositions
}
