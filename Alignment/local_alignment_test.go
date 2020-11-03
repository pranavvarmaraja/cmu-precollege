package Alignment

import (
	"reflect"
	"testing"
)

type Solution struct {
	score  float64
	start0 int
	end0   int
	start1 int
	end1   int
}

type testpair struct {
	input LocalAlignmentInput
	sol   Solution
}

var tests = []testpair{
	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		Solution{2.000, 1, 3, 1, 3}},

	{LocalAlignmentInput{"TAAC",
		"TAAG",
		1.0, 1.0, 0.5},
		Solution{3.000, 0, 3, 0, 3}},

	{LocalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		Solution{1.000, 0, 1, 0, 1}},

	{LocalAlignmentInput{"TACGG",
		"CACGTG",
		1.0, 1.0, 0.5},
		Solution{4.500, 1, 5, 1, 6}},

	{LocalAlignmentInput{"AGATCGG",
		"AGCGGATTTGCC",
		1.0, 1.0, 0.5},
		Solution{6.000, 0, 7, 0, 5}},

	{LocalAlignmentInput{"AGATCGG",
		"AGTTGATTTGCCG",
		1.0, 0.5, 1.0},
		Solution{3.000, 1, 4, 4, 7}},

	{LocalAlignmentInput{"AGATCGG",
		"AGTCAGCGG",
		1.0, 1.0, 0.5},
		Solution{7.500, 0, 7, 0, 8}},

	{LocalAlignmentInput{"GCGATCGAT",
		"GCGCGATCTTGAT",
		1.0, 1.0, 0.5},
		Solution{10.000, 0, 9, 2, 13}}}

func TestLocalAlignment(t *testing.T) {
	for _, pair := range tests {
		optAlignment, start0, end0, start1, end1 := LocalAlignment(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		score := computeScore(optAlignment, pair.input.match, pair.input.mismatch, pair.input.gap)
		v := Solution{score, start0, end0, start1, end1}
		if !reflect.DeepEqual(v, pair.sol) {
			t.Error(
				"For", pair.input,
				"expected alignment with score", pair.sol.score,
				"got", optAlignment, "with score", score,
			)
		}
	}
}
