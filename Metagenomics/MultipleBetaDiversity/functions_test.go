package MultipleBetaDiversity

import (
	BrayCurtis "Metagenomics/BrayCurtis"
	Jaccard "Metagenomics/Jaccard"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type testpair struct {
	filename    string
	sampleNames []string
	allMaps     [](map[string]int)
	distMetric  string
	distMatrix  [][]float64
}

var f1 = "test1"
var f2 = "test2"
var f3 = "test3"
var f4 = "test4"
var f5 = "test5"
var f6 = "test6"
var f7 = "test7"
var f8 = "test8"
var f9 = "test9"
var f10 = "test10"
var f11 = "test11"
var f12 = "test12"

var f1map = FreqMapFromFile(filepath.Join("Tests", "test1.txt"))
var f2map = FreqMapFromFile(filepath.Join("Tests", "test2.txt"))
var f3map = FreqMapFromFile(filepath.Join("Tests", "test3.txt"))
var f4map = FreqMapFromFile(filepath.Join("Tests", "test4.txt"))
var f5map = FreqMapFromFile(filepath.Join("Tests", "test5.txt"))
var f6map = FreqMapFromFile(filepath.Join("Tests", "test6.txt"))
var f7map = FreqMapFromFile(filepath.Join("Tests", "test7.txt"))
var f8map = FreqMapFromFile(filepath.Join("Tests", "test8.txt"))
var f9map = FreqMapFromFile(filepath.Join("Tests", "test9.txt"))
var f10map = FreqMapFromFile(filepath.Join("Tests", "test10.txt"))
var f11map = FreqMapFromFile(filepath.Join("Tests", "test11.txt"))
var f12map = FreqMapFromFile(filepath.Join("Tests", "test12.txt"))

func makeMatrix(dim int) [][]float64 {
	var M = make([][]float64, dim)

	for i := range M {
		M[i] = make([]float64, dim)
	}
	return M
}

var tests = []testpair{
	{filepath.Join("Tests", "BrayCurtisMatrix.txt"),
		[]string{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12},
		[](map[string]int){f1map, f2map, f3map, f4map, f5map, f6map, f7map, f8map, f9map, f10map, f11map, f12map},
		"Bray-Curtis",
		makeMatrix(12)},
	{filepath.Join("Tests", "JaccardMatrix.txt"), []string{f7, f8, f9, f10, f11, f12},
		[](map[string]int){f7map, f8map, f9map, f10map, f11map, f12map},
		"Jaccard",
		makeMatrix(6)}}

func TestBetaDiversityMatrix(t *testing.T) {
	for _, pair := range tests {
		v := BetaDiversityMatrix(pair.allMaps, pair.distMetric)
		var d float64
		for i := range pair.sampleNames {
			for j := range pair.sampleNames {
				switch pair.distMetric {
				case "Bray-Curtis":
					d = BrayCurtis.BrayCurtisDistance(pair.allMaps[i], pair.allMaps[j])
				case "Jaccard":
					d = Jaccard.JaccardDistance(pair.allMaps[j], pair.allMaps[i])
				}
				pair.distMatrix[i][j] = d
			}
		}
		if !reflect.DeepEqual(v, pair.distMatrix) {
			t.Error(
				"For", pair.allMaps,
				"with dist metric", pair.distMetric,
				"expected", pair.distMatrix,
				"got", v,
			)
		}
		WriteMatrixToFile(pair.distMatrix, pair.sampleNames, pair.distMetric, pair.filename)
	}
}

func FreqMapFromFile(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")
	freqMap := make(map[string]int, 0)
	for _, line := range lines {
		l := strings.Split(line, ":")
		key, val := l[0], strings.TrimSpace(l[1])
		freqMap[key], _ = strconv.Atoi(val)
	}
	return freqMap
}

func WriteMatrixToFile(matrix [][]float64, sampleNames []string, distMetric string, filename string) {
	res := "\t"

	for _, name := range sampleNames {
		res = res + name + "\t"
	}

	for i, name := range sampleNames {
		res = res + "\n" + name + "\t"
		for j, _ := range sampleNames {
			res = res + strconv.FormatFloat(matrix[i][j], 'f', 3, 64) + "\t"
		}
	}
	ioutil.WriteFile(filename, []byte(res), 0644)
}
