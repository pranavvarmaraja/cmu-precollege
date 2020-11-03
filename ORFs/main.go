package main

import (
	"fmt"
)

func main() {
	fmt.Println("ORF finding.")

	filename := "Data/rona_genome.fasta"

	genome := ReadGenomeFASTA(filename)

	rnaStrand := DNAToRNA(genome) //transcribe

	minORFLength := 300

	//find ORFs in rna strand
	orfs0 := FindORFsRNA2(rnaStrand, minORFLength, false)

	// iterate this process on the opposing strand

	genomeRC := ReverseComplement(genome)

	//convert to RNA
	rnaStrandRC := DNAToRNA(genomeRC)

	//find ORFs on opposing strand
	orfs1 := FindORFsRNA2(rnaStrandRC, minORFLength, true)

	orfs := append(orfs0, orfs1...) // final set of ORFs (candidate genes)

	//let's write them to a file

	outFilename := "Output/rona_ORFs.txt"

	WriteORFsToFile(orfs, outFilename)

}
