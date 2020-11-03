package main

//ORF is an object that contains the starting position and length of the ORF, along with whether it occurs on the revComp of the genome (true) or the main strand (false)
type ORF struct {
	startingPosition, length int // length in # nucleotides
	revComp                  bool
}

// Implement FindORFsRNA using your Translate function.  This should take as
// input an RNA string rnaStrand and an integer minORFLength and a boolean rc
// indicating whether we are looking at the reverse complement or not.  This
// function should return all of the ORFs discovered.

func FindORFsRNA(rnaStrand string, minORFLength int, rc bool) []ORF {
	if minORFLength <= 0 {
		panic("Error: minimum ORF length is nonpositive in FindORFsRNA function.")
	}

	orfs := make([]ORF, 0) // every time we find an ORF we will append it

	//ranging through every possible start index and adding any ORF we find to our slice
	for startIndex := 0; startIndex < len(rnaStrand)-minORFLength+1-3; startIndex++ {
		readingFrame := Translate(rnaStrand, startIndex)
		//most of the time we get an empty string
		//how long was ORF found?
		orfLength := len(readingFrame) * 3
		if orfLength >= minORFLength {
			//this assumes that minORFLength is positive which we tested above
			//create my ORF variable
			var o ORF
			//set its fields
			o.startingPosition = startIndex
			o.length = orfLength
			o.revComp = rc

			// now append our new ORF to our growing list
			orfs = append(orfs, o)
		}
	}

	return orfs
}
