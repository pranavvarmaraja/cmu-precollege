package main

import "fmt"

func main() {
	fmt.Println("Genome assembly!")

	//fmt.Println("First, greedy assembly.")

	//RunGreedyAssembly()

	fmt.Println("Let's build an overlap network for a simulated dataset.")

	BuildOverlapNetwork()
}

func BuildOverlapNetwork() {
	genomeLength := 2000
	readLength := 200

	//first, simulate a genome
	genome := GenerateRandomGenome(genomeLength)

	//establish a probability of picking a given position
	probability := 0.1

	//establish error rates
	insertionErrorRate := 0.003
	deletionErrorRate := 0.003
	substitutionErrorRate := 0.003

	//simulate some messy reads
	reads := SimulateReadsMessy(genome, readLength, probability, substitutionErrorRate, insertionErrorRate, deletionErrorRate)

	fmt.Println("Reads simulated. We have", len(reads), "total reads.")
	fmt.Println("Making overlap network.")

	//need sequence alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//need threshold parameter for binarizing the overlap network
	threshold := 50.0

	//form overlap network.
	adjList := MakeOverlapNetwork(reads, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(adjList), "total keys.")

	fmt.Println("The average outdegree of a node is", AverageOutDegree(adjList))

	fmt.Println("Now let's clean up the overlap network.")

	maxK := 3

	adjList = TrimNetwork(adjList, maxK)

	fmt.Println("Overlap network trimmed! Now the network has", len(adjList), "total keys.")

	fmt.Println("The average outdegree of a node is", AverageOutDegree(adjList))

	contigs := GenerateContigs(adjList)
	//contigs has type [][]string

	fmt.Println("Contigs generated! We have", len(contigs), "total contigs.")

	fmt.Println("Now we will generate a multiple alignment for each contig.")

	//sequencedContigs will hold our actual strings
	sequencedContigs := make([]string, len(contigs))

	//what is my threshold for what frequency of symbols is good enough to report a winner?
	voteThreshold := 0.5

	for i := range contigs {
		columnCounts := MultAlignmentFromContig(contigs[i])
		// we need to infer the string from columnCounts
		// take a vote!
		sequencedContigs[i] = ConsensusFromMultAlign(columnCounts, voteThreshold)
	}

	fmt.Println("We made it! Contigs converted to strings.")

	// how can we determine if our inferred genome is close to the original?
	// let's write our genome and our contigs to a file
	genomeOutFile := "Output/genome.fasta"
	contigsOutFile := "Output/contigs.fasta"

	WriteGenomeToFileFASTA(genome, genomeOutFile)
	WriteContigsToFileFASTA(sequencedContigs, contigsOutFile)
}

func RunGreedyAssembly() {
	genomeLength := 100
	k := 10
	//first, generate a random genome

	originalGenome := GenerateRandomGenome(genomeLength)

	//form its k-mer composition
	reads := KmerComposition(originalGenome, k)

	//now, call our greedy assembler
	assembledGenome := GreedyAssembler(reads)

	kmers := KmerComposition(assembledGenome, k)

	if StrSliceEq(reads, kmers) == true {
		fmt.Println("Mission accomplished!")
	}
}
