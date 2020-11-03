package Alignment

import (
	"reflect"
	"testing"
)

type levenshteinMatrixTestpair struct {
	patterns []string
	matrix   [][]int
}

var levenshteinMatrixTests = []levenshteinMatrixTestpair{
	{[]string{"A", "T"}, [][]int{[]int{0, 1}, []int{1, 0}}},
	{[]string{"-GCTAG--TACCTATCGGA---", "TAGCTAG---TCGAT", "AGC---TAGGGATCGAAAT----"},
		[][]int{[]int{0, 12, 11}, []int{12, 0, 13}, []int{11, 13, 0}}},
	{[]string{"ATCGCTAGCTCTTCGATC", "TTTCGATCTTAAGATC"}, [][]int{[]int{0, 8}, []int{8, 0}}},
	{[]string{"ATCGCTAGCTCTTCGATC", "TTTCGATCTTAAGATC", "--TTTCGATC---TTAAGATC--", "ATC--GCTAG-CTCTTCG---ATC"},
		[][]int{[]int{0, 8, 13, 6}, []int{8, 0, 7, 14}, []int{13, 7, 0, 17}, []int{6, 14, 17, 0}}},
	{[]string{"A", "A", "A"}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}}}

func TestEditDistanceMatrix(t *testing.T) {
	for _, pair := range levenshteinMatrixTests {
		v := EditDistanceMatrix(pair.patterns)
		if !reflect.DeepEqual(v, pair.matrix) {
			t.Error(
				"For", pair.patterns,
				"expected", pair.matrix,
				"got", v,
			)
		}
	}
}
