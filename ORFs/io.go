package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadGenomeFASTA(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: something went wrong opening the file.")
		fmt.Println("Probably you gave the wrong filename.")
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine[0] != '>' { // FASTA header
			lines = append(lines, currentLine)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: there was a problem reading the file.")
		os.Exit(1)
	}

	if scanner.Err() != nil {
		panic("Error: there was a problem reading the file.")
	}

	file.Close()

	genome := ""

	for _, line := range lines {
		genome += line
	}

	if ValidDNA(genome) == false {
		panic("Genome passed into ReadGenomeFASTA is not a valid DNA string.")
	}

	return genome
}

//ValidDNA returns true if a given string only has the symbols "A", "C", "G", and "T". It returns false otherwise.
func ValidDNA(read string) bool {
	for _, symbol := range read {
		if symbol != 'A' && symbol != 'C' && symbol != 'G' && symbol != 'T' {
			return false
		}
	}
	return true
}
