package main

import "testing"

func StringIndexEq(index1, index2 StringIndex) bool {
	if len(index1) != len(index2) {
		return false
	}
	for i := range index1 {
		if !IntSliceEq(index1[i], index2[i]) {
			return false
		}
	}
	return true
}

/* TESTS */

//Testing Updated Version of MapToMinimizer
func TestMapToMinimizer(t *testing.T) {
	reads1 := []string{
		"ACCGATC",
		"AGGTGG",
		"CAGAGCC",
		"GGACGTA",
		"TTCGAGG",
		"TTGACCA",
	}
	reads2 := []string{
		"AAGGCCAGGGA",
		"AGGTCGAGGCGT",
		"AGTCCGAGTTCCA",
		"ATTCGACGAA",
		"CATTCGGGTAA",
		"CCAGTAATTGA",
		"GGGTAGGACAGGG",
		"TTAGGCGAGTC",
		"TTGGACGAGCC",
	}

	ans1 := StringIndex{
		"AC": {0, 3, 5},
		"AG": {1, 2, 4},
		"AT": {0},
		"CA": {2},
		"CG": {4},
		"GA": {5},
		"GG": {1, 3},
		"TT": {4, 5},
	}
	ans2 := StringIndex{
		"ACC": {0, 5},
		"ACG": {3},
		"AGA": {2},
		"AGG": {1, 4},
		"CAG": {2},
		"GGA": {3},
		"TTC": {4},
		"TTG": {5},
	}
	ans3 := StringIndex{
		"AAGGC": {0},
		"AATTG": {5},
		"ACAGG": {6},
		"ACGAA": {3},
		"ACGAG": {8},
		"AGGCG": {1, 7},
		"AGTCC": {2},
		"ATTCG": {4},
	}

	//testcase#1
	result1 := MapToMinimizer(reads1, 2, 3)
	if !StringIndexEq(result1, ans1) {
		t.Error("MapToMinimizer failed testcase1", "Failed on:", reads1, 2, 3,
			"\n", "Expecting:", ans1, "Actual result:", result1)
	}

	//testcase#2
	result2 := MapToMinimizer(reads1, 3, 2)
	if !StringIndexEq(result2, ans2) {
		t.Error("MapToMinimizer failed testcase2", "Failed on:", reads1, 3, 2,
			"\n", "Expecting:", ans2, "Actual result:", result2)
	}

	//testcase#3
	result3 := MapToMinimizer(reads2, 5, 1)
	if !StringIndexEq(result3, ans3) {
		t.Error("MapToMinimizer failed testcase3", "Failed on:", reads2, 5, 1,
			"\n", "Expecting:", ans3, "Actual result:", result3)
	}
}
