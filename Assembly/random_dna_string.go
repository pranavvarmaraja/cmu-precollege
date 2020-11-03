package main

import (
	"math/rand"
)

//GenerateRandomGenome takes a parameter genomeLength and returns
//a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(genomeLength int) string {
	// make an array of symbols of the appropriate length
	symbols := make([]byte, genomeLength)
	for i := 0; i < genomeLength; i++ {
		symbols[i] = RandomDnaSymbol()
	}
	return string(symbols)
}

//RandomDnaSymbol produces 'A', 'C', 'G', or 'T' with equal probability.
func RandomDnaSymbol() byte {
	i := rand.Intn(4) // integer btw 0 and 3 inclusively
	symbols := []byte{'A', 'C', 'G', 'T'}
	return symbols[i]
}
