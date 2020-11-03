package main

import "math/rand"

//SimulateReadsClean takes a genome along with a read length and a probability.
//It returns a collection of strings resulting from simulating clean reads,
//where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	n := len(genome)
	reads := make([]string, 0)

	for i := 0; i < n-readLength+1; i++ {
		x := rand.Float64()
		if x < probability {
			//add the k-mer starting at position i to our read set
			reads = append(reads, genome[i:i+readLength])
		}
	}

	return reads
}

//NOTE: we will simulate messy reads together.

//SimulateReadsMessy takes a genome along with a read length, a probability,
//and error rates for substitutions, insertions, and deletions.
//It returns a collection of reads resulting from simulating clean reads,
//where a given position is sampled with the given probability, and
//errors occur at the respective rates.
func SimulateReadsMessy(genome string, readLength int, probability float64, substitutionErrorRate, insertionErrorRate, deletionErrorRate float64) []string {
	reads := make([]string, 0)

	if substitutionErrorRate+insertionErrorRate+deletionErrorRate >= 1 {
		panic("Error: Sean is a stickler :)")
	}

	//simulate our messy reads along the way

	//sample starting positions

	for i := 0; i < len(genome)-readLength+1; i++ {
		//roll our die to see if we pick a read at this location
		x := rand.Float64()
		if x < probability {
			//build a string
			currString := ""
			currPosition := i
			//continue until we get a read of the appropriate length
			for len(currString) < readLength && currPosition < len(genome) { // second test assures currPosition doesn't go out of range and give us an out of bounds error
				//roll another die to determine one of four outcomes: take current symbol, substitute, insert, or delete
				y := rand.Float64()
				if y < substitutionErrorRate {
					//substitute!
					sym := RandomDnaSymbol()
					// possible problem: if symbol generated is the same :(
					// solution: keep generating the symbol if it's different
					for sym == genome[currPosition] {
						sym = RandomDnaSymbol()
					}
					// now I can guarantee that sym is different from the current position
					currString += string(sym)
					currPosition++
				} else if y < substitutionErrorRate+insertionErrorRate {
					currString += string(RandomDnaSymbol())
					//don't need to increase the current position
				} else if y < substitutionErrorRate+insertionErrorRate+deletionErrorRate {
					// delete!
					//think of this as "skip over" current position in the genome
					currPosition++
				} else {
					// take current symbol
					currString += string(genome[currPosition])
					currPosition++
				}
			}
			//the string is built BUT there's a small chance it is too short if we hit the end of the genome
			if len(currString) == readLength {
				reads = append(reads, currString)
			}
		}
	}

	return reads
}
