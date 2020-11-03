package Alignment

import (
	"reflect"
	"testing"
)

type LocalAlignmentInput struct {
	str1     string
	str2     string
	match    float64
	mismatch float64
	gap      float64
}

type localscoretestpair struct {
	input    LocalAlignmentInput
	scoreMatrix [][]float64
}

var localScoreTests = []localscoretestpair{
	{LocalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0,0}, []float64{0,1}}},

	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 0, 0, 1},
					[]float64{0, 0, 1, 1, 0.5},
					[]float64{0, 0, 1, 2, 1.5},
					[]float64{0, 1, 0.5, 1.5, 1}}},

	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 0, 0, 1},
					[]float64{0, 0, 1, 1, 0},
					[]float64{0, 0, 1, 2, 1},
					[]float64{0, 1, 0, 1, 1.5}}},

	{LocalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 1, 0.5, 0},
					[]float64{0, 1, 0.5, 2, 1.5},
					[]float64{0, 0.5, 0, 1.5, 3}}},

	{LocalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 0, 0, 0},
					[]float64{0, 0, 1, 0, 0},
					[]float64{0, 1, 0, 2, 1},
					[]float64{0, 0, 0.5, 1, 3},
					[]float64{0, 0, 0, 0, 2}}},

	{LocalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					[]float64{0, 1, 0.5, 0, 0, 0, 0, 0, 1, 0.5},
					[]float64{0, 0.5, 2, 1.5, 1, 0.5, 0, 1, 0.5, 0},
					[]float64{0, 0, 1.5, 3, 2.5, 2, 1.5, 1, 0.5, 1.5},
					[]float64{0, 0, 1, 2.5, 4, 3.5, 3, 2.5, 2, 1.5},
					[]float64{0, 1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5, 3},
					[]float64{0, 0.5, 2, 1.5, 3, 2.5, 2, 3.5, 3, 2.5},
					[]float64{0, 0, 1.5, 3, 2.5, 2, 3.5, 3, 2.5, 4},
					[]float64{0, 0, 1, 2.5, 4, 3.5, 3, 2.5, 2, 3.5},
					[]float64{0, 0, 1, 2, 3.5, 3, 2.5, 4, 3.5, 3}}},

	{LocalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0},
					[]float64{0, 1, 1, 0.5},
					[]float64{0, 0.5, 0.5, 0},
					[]float64{0, 0, 0, 1.5},
					[]float64{0, 0, 0, 1},
					[]float64{0, 1, 1, 0.5},
					[]float64{0, 0.5, 0.5, 0},
					[]float64{0, 0, 0, 1.5},
					[]float64{0, 0, 0, 1},
					[]float64{0, 0, 0, 0.5}}}}

func TestLocalScoreTable(t *testing.T) {
	for _, pair := range localScoreTests {
		v := LocalScoreTable(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		if !reflect.DeepEqual(v, pair.scoreMatrix) {
			t.Error(
				"For", pair.input,
				"expected scoring matrix", pair.scoreMatrix,
				"got", v,
			)
		}
	}
}
