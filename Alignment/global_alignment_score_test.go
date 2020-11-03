package Alignment

import (
	"reflect"
	"testing"
)

type GlobalAlignmentInput struct {
	str1     string
	str2     string
	match    float64
	mismatch float64
	gap      float64
}

type globaltestpair struct {
	input    GlobalAlignmentInput
	solScore float64
}

type globalscoretestpair struct {
	input       GlobalAlignmentInput
	scoreMatrix [][]float64
}

var globalScoreTests = []globalscoretestpair{
	{GlobalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1}, []float64{-1, 1}}},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2},
			[]float64{-0.5, -1, -1.5, -2, -0.5},
			[]float64{-1, -1.5, 0, -0.5, -1},
			[]float64{-1.5, -2, -0.5, 1, 0.5},
			[]float64{-2, -0.5, -1, 0.5, 0}}},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1, -2, -3, -4},
			[]float64{-1, -0.5, -1.5, -2.5, -2},
			[]float64{-2, -1.5, 0.5, -0.5, -1.5},
			[]float64{-3, -2.5, -0.5, 1.5, 0.5},
			[]float64{-4, -2, -1.5, 0.5, 1}}},

	{GlobalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2},
			[]float64{-0.5, -1, -1.5, -2, -2.5},
			[]float64{-1, -1.5, 0, -0.5, -1},
			[]float64{-1.5, 0, -0.5, 1, 0.5},
			[]float64{-2, -0.5, -1, 0.5, 2}}},

	{GlobalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1, -2, -3, -4},
			[]float64{-1, -0.5, -1.5, -2.5, -3.5},
			[]float64{-2, -1.5, 0.5, -0.5, -1.5},
			[]float64{-3, -1, -0.5, 1.5, 0.5},
			[]float64{-4, -2, -1.5, 0.5, 2.5},
			[]float64{-5, -3, -2.5, -0.5, 1.5}}},

	{GlobalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2, -2.5, -3, -3.5, -4, -4.5},
			[]float64{-0.5, 1, 0.5, 0, -0.5, -1, -1.5, -2, -2.5, -3},
			[]float64{-1, 0.5, 2, 1.5, 1, 0.5, 0, -0.5, -1, -1.5},
			[]float64{-1.5, 0, 1.5, 3, 2.5, 2, 1.5, 1, 0.5, 0},
			[]float64{-2, -0.5, 1, 2.5, 4, 3.5, 3, 2.5, 2, 1.5},
			[]float64{-2.5, -1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5, 3},
			[]float64{-3, -1.5, 0, 1.5, 3, 2.5, 2, 3.5, 3, 2.5},
			[]float64{-3.5, -2, -0.5, 1, 2.5, 2, 3.5, 3, 2.5, 4},
			[]float64{-4, -2.5, -1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5},
			[]float64{-4.5, -3, -1.5, 0, 1.5, 3, 2.5, 4, 3.5, 3}}},

	{GlobalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5},
			[]float64{-0.5, 1, 0.5, 0},
			[]float64{-1, 0.5, 0, -0.5},
			[]float64{-1.5, 0, -0.5, 1},
			[]float64{-2, -0.5, -1, 0.5},
			[]float64{-2.5, -1, 0.5, 0},
			[]float64{-3, -1.5, 0, -0.5},
			[]float64{-3.5, -2, -0.5, 1},
			[]float64{-4, -2.5, -1, 0.5},
			[]float64{-4.5, -3, -1.5, 0}}}}

func computeScore(alignment [2]string, match float64, mismatch float64, gap float64) float64 {
	score := 0.0
	str1 := alignment[0]
	str2 := alignment[1]
	gapC := "-"

	for i, c := range str1 {
		if string(c) == string(str2[i]) {
			score += match
		} else if string(c) == gapC || string(str2[i]) == gapC {
			score += gap
		} else {
			score += mismatch
		}
	}
	return score
}

func TestGlobalScoreTable(t *testing.T) {
	for _, pair := range globalScoreTests {
		v := GlobalScoreTable(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		if !reflect.DeepEqual(v, pair.scoreMatrix) {
			t.Error(
				"For", pair.input,
				"expected scoring matrix", pair.scoreMatrix,
				"got", v,
			)
		}
	}
}
