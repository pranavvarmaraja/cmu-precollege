package Alignment

import (
	"testing"
)

type longestSharedSubstringTestpair struct {
	str1      string
	str2      string
	substring []string
}

var longestSharedSubstringTests = []longestSharedSubstringTestpair{
	{"ATGCGGCTAGCTTAGCCTAGATCGATCGGCTAGCTAGCTAGCCGAGGCTCTCGATCGATCGCGCTAGG",
		"ATGCGGCTGGCTTAGCCTAGTTCGATCGCGTTCGTAGCTATAGAGCTAGCTAGATCGATCGCGCTAGG",
		[]string{"GATCGATCGCGCTAGG"}},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"GGTACATCGATTCTAGATTCTATAGCGCGCTTCGATCGATTCGATCGATCGAAAAG",
		[]string{"ATTCGATCG"}},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		[]string{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT"}},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATAGCATCGGGAACACAGATCGAT",
		[]string{"TTGCGGAGCTAGGGATCCGATCGAATATCGATAT"}},
	{"AGCT",
		"TCGA",
		[]string{"A", "T", "G", "C"}},
	{"AGCTGAT",
		"AGCCGAT",
		[]string{"AGC", "GAT"}},
	{"AATT",
		"CCGG",
		[]string{""}},
	{"AGCT",
		"TGAC",
		[]string{"A", "T", "G", "C"}}}

func TestLongestSharedSubstring(t *testing.T) {
	for _, pair := range longestSharedSubstringTests {
		v := LongestSharedSubstring(pair.str1, pair.str2)
		notFound := true
		for _, val := range pair.substring {
			if v == val {
				notFound = false
			}
		}
		if notFound {
			t.Error(
				"For", pair.str1,
				"and", pair.str2,
				"expected one of", pair.substring,
				"got", v,
			)
		}
	}
}
