package main

//import (
//	"sort"
//)

//ORF is an object that contains the starting position and length of the ORF, along with whether it occurs on the revComp of the genome (true) or the main strand (false)

// Implement FindORFsRNA using your Translate function.  This should take as
// input an RNA string rnaStrand and an integer minORFLength and a boolean rc
// indicating whether we are looking at the reverse complement or not.  This
// function should return the ORF of maximal and acceptable length from
// each start codon.

func FindORFsRNA2(rnaStrand string, minORFLength int, rc bool) []ORF {
	if minORFLength <= 0 {
		panic("Error: minimum ORF length is nonpositive in FindORFsRNA function.")
	}

	orfs := make([]ORF, 0) // every time we find an ORF we will append it

	//let's make a map to see if we have encountered the index of a given STOP codon before
	stopCodons := make(map[int]int) //keys are positions of stop codons, values are position of earliest start codon for this ORF

	//ranging through every possible start index and adding any ORF we find to our slice
	for startIndex := 0; startIndex < len(rnaStrand)-minORFLength+1-3; startIndex++ {
		readingFrame := Translate(rnaStrand, startIndex)
		//most of the time we get an empty string
		//how long was ORF found?
		orfLength := len(readingFrame) * 3
		if orfLength >= minORFLength {
			//difference from previous function is that we only want to create an ORF
			//if we haven't seen the current stop codon before

			// we will use the stopCodons map to ask, "have I seen this stop codon before?"
			stopIndex := startIndex + orfLength
			_, exists := stopCodons[stopIndex] // if exists == false this is a new ORF
			if exists == false {
				//create an entry in the stopIndex map for this start index
				stopCodons[stopIndex] = startIndex

				//also create my ORF here

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
	}

	return orfs
}
