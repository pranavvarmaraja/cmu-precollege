package main

//GreedyAssembler takes a collection of strings and returns a genome whose
//k-mer composition is these strings. It makes the following assumptions.
//1. "Perfect coverage" -- every k-mer is detected
//2. No errors in reads
//3. Every read has equal length (k)
//4. DNA is single-stranded
func GreedyAssembler(reads []string) string {
	//create a copy of our reads that we can edit
	reads2 := make([]string, len(reads))
	copy(reads2, reads)

	if len(reads) == 0 {
		panic("Error: no reads given to Greedy assembler.")
	}

	k := len(reads[0])

	genome := reads2[0]

	//delete 0-th value from our slice
	reads2 = append(reads2[:0], reads2[1:]...)

	//while we still have reads, try to extend current read
	for len(reads2) > 0 {
		for i, kmer := range reads2 {
			//try to extend genome to left and right
			//a hit means that we match k-1 nucleotides
			if genome[0:k-1] == kmer[1:] {
				//append to the left side of genome
				genome = kmer[0:1] + genome
				reads2 = append(reads2[:i], reads2[i+1:]...) // throw out current read
				//stop current test
				break
			} else if genome[len(genome)-k+1:] == kmer[:k-1] {
				//append to the right side of genome
				genome = genome + kmer[k-1:]
				reads2 = append(reads2[:i], reads2[i+1:]...) // throw out current read
				break
			}
		}
	}
	return genome
}
