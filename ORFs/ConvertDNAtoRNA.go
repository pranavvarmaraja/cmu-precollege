package main

//Implement a function called DNAtoRNA which takes as input a DNA string genome and returns the RNA version of that string.

func DNAToRNA(genome string) string {
	rna := make([]rune, len(genome))

	//range over the genome and set the appropriate symbol of our slice
	for i, symbol := range genome {
		if symbol == 'A' || symbol == 'C' || symbol == 'G' {
			rna[i] = symbol
		} else if symbol == 'T' {
			//thymine replaced with uracil
			rna[i] = 'U'
		} else {
			panic("Error: Non-DNA string passed to DNAToRNA.")
		}
	}

	//convert rna to a string
	return string(rna)
}
