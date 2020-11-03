package Alignment

import (
	"strconv"
	"testing"
)

var globalTests = []globaltestpair{
	{GlobalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		1.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		3.000},

	{GlobalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		4.500},

	{GlobalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		9.000},

	{GlobalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		6.000}}

func TestGlobalAlignment(t *testing.T) {
	for _, pair := range globalTests {
		v := GlobalAlignment(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		score := computeScore(v, pair.input.match, pair.input.mismatch, pair.input.gap)
		if score != pair.solScore {
			t.Error(
				"For", pair.input,
				"expected alignment with score", strconv.FormatFloat(pair.solScore, 'f', 3, 64),
				"got", v, "with a score of", strconv.FormatFloat(score, 'f', 3, 64),
			)
		}
	}
}
