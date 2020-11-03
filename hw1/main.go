package main

import (
	"fmt"
)

func main() {

	genome := "GATCAGCATAAGGGTCCCTGCAATGCATGACAAGCCTGCAGTTGTTTTAC"
	fmt.Println(FindClumps(genome, 4, 25, 3))

}

func Reverse(s string) string {
	a := ""
	for i := len(s) - 1; i >= 0; i-- {

		a += string(s[i])
	}
	return a
}

func Complement(dna string) string {

	ret := ""
	for _, val := range dna {
		switch val {
		case 'A':
			ret += "T"
		case 'T':
			ret += "A"
		case 'C':
			ret += "G"
		case 'G':
			ret += "C"
		}
	}
	return ret
}

func ReverseComplement(dna string) string {
	return Reverse(Complement(dna))
}

func PatternCount(pattern string, text string) int {

	indices := StartingIndices(pattern, text)
	return len(indices)
}

func StartingIndices(pattern string, text string) []int {

	ret := make([]int, 0)

	for i := 0; i < len(text)-len(pattern)+1; i++ {
		if text[i:i+len(pattern)] == pattern {
			ret = append(ret, i)
		}
	}
	return ret
}

func FrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)
	freqMap := FrequencyTable(text, k)
	max := MaxMap(freqMap)

	for key, value := range freqMap {
		if value == max {
			freqPatterns = append(freqPatterns, key)
		}
	}
	return freqPatterns
}

func FrequencyTable(text string, k int) map[string]int {
	ret := make(map[string]int)
	for i := 0; i < len(text)-k+1; i++ {
		if ret[text[i:i+k]] == 0 {
			ret[text[i:i+k]] = 1
		} else {
			ret[text[i:i+k]]++
		}
	}
	return ret
}

func MaxMap(freqmap map[string]int) int {
	ret := 0
	firstTime := true
	for _, val := range freqmap {
		if firstTime == true || ret < val {
			ret = val
			firstTime = false
		}
	}
	return ret
}

func FindClumps(text string, k int, L int, t int) []string {
	patterns := make([]string, 0)
	for i := 0; i < len(text)-L+1; i++ {
		window := text[i : i+L]
		freqMap := FrequencyTable(window, k)
		for key, _ := range freqMap {
			if freqMap[key] >= t && !Contains(patterns, key) {
				patterns = append(patterns, key)
			}
		}
	}
	return patterns
}

func Contains(arr []string, s string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func SkewArray(genome string) []int {

	ret := make([]int, len(genome)+1)

	for i := 1; i < len(ret); i++ {
		ret[i] = ret[i-1] + Skew(genome[i-1])
	}
	return ret
}

func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	}
	return 0
}

func MinimumSkew(genome string) []int {
	indices := make([]int, 0)
	skewArr := SkewArray(genome)
	min := MinValue(skewArr)
	for i, val := range skewArr {
		if val == min {
			indices = append(indices, i)
		}
	}
	return indices
}

func MinValue(arr []int) int {
	min := arr[0]

	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}
