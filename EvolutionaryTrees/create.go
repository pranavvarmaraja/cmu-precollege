package main

import (
	// "fmt"
	"strings"
	// "EvolutionaryTrees/Lib"
	"EvolutionaryTrees/Clustal"
	"EvolutionaryTrees/Io"
	"EvolutionaryTrees/Lib"
	"EvolutionaryTrees/UPGMA"
	// "EvolutionaryTrees/Clustal"
)

func main() {

	/*
	   // Create Garfield Tree
	   mtx, labels := Io.ReadMatrixFromFile("Datasets/Input/Garfield/distance.mtx")
	   T := UPGMA.UPGMA(mtx, labels) // CODE IS RUN
	   Io.WriteNewickToFile(T, "Datasets/Output/Garfield", "garfield.tre")
	*/

	// Create Garfield Alignment
	// dnaMap := Io.ReadDNAStringsFromFile("Datasets/Input/Garfield/RAW/garfield.txt")
	// newLabels, newDnaStrings := Lib.GetKeyValues(dnaMap)
	// dnaStrings := rearrangeStrings(labels, newLabels, newDnaStrings)
	// alignt := Clustal.MultiAlign(T, dnaStrings)
	// Io.WriteAlignmentToFile(alignt, Io.SequenceOrder(T), "Datasets/Output/Garfield", "garfield.fa")

	// Create Example Tree
	mtx, labels := Io.ReadMatrixFromFile("Datasets/Input/Test-Example/distance.mtx")
	T := UPGMA.UPGMA(mtx, labels) // CODE IS RUN
	Io.WriteNewickToFile(T, "Datasets/Output/Test-Example", "toy.tre")

	//set alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 6.0
	supergap := 10.0

	// Create Example Alignment
	dnaMap := Io.ReadDNAStringsFromFile("Datasets/Input/Test-Example/RAW/toy-example.fasta")
	newLabels, newDnaStrings := Lib.GetKeyValues(dnaMap)
	dnaStrings := rearrangeStrings(labels, newLabels, newDnaStrings)
	alignt := Clustal.MultiAlign(T, dnaStrings, match, mismatch, gap, supergap)
	Io.WriteAlignmentToFile(alignt, Io.SequenceOrder(T), "Datasets/Output/Test-Example", "toy.fa")

	/*
	   // Create Coronavirus Tree
	   cats := []string{"USA", "CHN", "ITA"}
	   mtx, labels = Io.ReadMatrixFromFile("Datasets/Input/SARS-Cov-2/distance.mtx")
	   T = UPGMA.UPGMA(mtx, labels) // CODE IS RUN
	   Io.WriteCSVToFile(T, cats, "Datasets/Output/SARS-Cov-2", "covid.csv")
	   Io.WriteNewickToFile(T, "Datasets/Output/SARS-Cov-2", "covid.tre")

	   // Create Coronavirus Alignment
	   // dnaMap := Io.ReadDNAStringsFromFile("Datasets/Input/SARS-Cov-2/RAW/covid-19-2.fasta")
	   // newLabels, newDnaStrings := Lib.GetKeyValues(dnaMap)
	   // dnaStrings := rearrangeStrings(labels, newLabels, newDnaStrings)
	   // alignt := Clustal.MultiAlign(T, dnaStrings)
	   // Io.WriteAlignmentToFile(alignt, "Datasets/Output/SARS-Cov-2", "covid.aln")

	   // Create Bacteria Tree
	   mtx, labels = Io.ReadMatrixFromFile("Datasets/Input/16S-Bact/distance.mtx")
	   cats = getCats(labels)
	   T = UPGMA.UPGMA(mtx, labels) // CODE IS RUN
	   Io.WriteCSVToFile(T, cats, "Datasets/Output/16S-Bact", "bact.csv")
	   Io.WriteNewickToFile(T, "Datasets/Output/16S-Bact", "bact.tre")

	   // Create Bacteria Alignment
	   // dnaMap := Io.ReadDNAStringsFromFile("Datasets/Input/16S-Bact/RAW/bacteriaF.txt")
	   // newLabels, newDnaStrings := Lib.GetKeyValues(dnaMap)
	   // dnaStrings := rearrangeStrings(labels, newLabels, newDnaStrings)
	   // alignt := Clustal.MultiAlign(T, dnaStrings)
	   // Io.WriteAlignmentToFile(alignt, Io.SequenceOrder(T), "Datasets/Output/16S-Bact", "bacteria.fa")
	*/
}

func getCats(labels []string) []string {
	cats := make([]string, 0)
	for _, label := range labels {
		cats = append(cats, strings.Split(label, "|")[0])
	}
	return unique(cats)
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func rearrangeStrings(newLabels []string, oldLabels []string, dnaStringsOld []string) []string {
	dnaStringsNew := make([]string, 0)
	for _, newLabel := range newLabels {
		j := getIndex(oldLabels, newLabel)
		dnaStringsNew = append(dnaStringsNew, dnaStringsOld[j])
	}
	return dnaStringsNew
}

func getIndex(arr []string, target string) int {
	for i, str := range arr {
		if str == target {
			return i
		}
	}
	return 0
}
