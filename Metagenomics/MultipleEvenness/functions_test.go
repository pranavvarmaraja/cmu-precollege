package MultipleEvenness

import ("strings"
		"testing"
		"path/filepath"
		"io/ioutil"
		"reflect"
		"strconv"
		"math"
		"Metagenomics/Evenness")

type testpair struct {
	filename string
	sampleNames []string
	allMaps [](map[string]int)
	simpsonArray []float64
}

var f1 = "test1"
var f2 = "test2"
var f3 = "test3"
var f4 = "test4"
var f5 = "test5"
var f6 = "test6"
var f7 = "test7"

var f1map = FreqMapFromFile(filepath.Join("Tests","test1.txt"))
var f2map = FreqMapFromFile(filepath.Join("Tests","test2.txt"))
var f3map = FreqMapFromFile(filepath.Join("Tests","test3.txt"))
var f4map = FreqMapFromFile(filepath.Join("Tests","test4.txt"))
var f5map = FreqMapFromFile(filepath.Join("Tests","test5.txt"))
var f6map = FreqMapFromFile(filepath.Join("Tests","test6.txt"))
var f7map = FreqMapFromFile(filepath.Join("Tests","test7.txt"))

var tests = []testpair{
{filepath.Join("Tests","SimpsonsMatrix.txt"), []string{f1,f2,f3,f4,f5,f6,f7},
 [](map[string]int){f1map,f2map,f3map,f4map,f5map,f6map,f7map},
 make([]float64, 7)}}

func TestSimpsonsMatrix(t *testing.T) {
	for _, pair := range tests {
		v := SimpsonsMatrix(pair.allMaps)
		for i := range pair.sampleNames {
			pair.simpsonArray[i] = Evenness.SimpsonsIndex(pair.allMaps[i])
		}
		if !reflect.DeepEqual(round(v), round(pair.simpsonArray)) {
			t.Error(
				"For", pair.allMaps,
				"expected", pair.simpsonArray,
				"got", v,
			)
		}
		WriteMatrixToFile(pair.simpsonArray, pair.sampleNames, pair.filename)
	}
}

func round(dist []float64) []float64 {
	for i, val := range dist {
		dist[i] = math.Round(val * float64(1000)) / float64(1000)
	}
	return dist
}

func FreqMapFromFile(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")
	freqMap := make(map[string]int, 0)
	for _,line := range lines {
		l := strings.Split(line, ":")
		key,val := l[0], strings.TrimSpace(l[1])
		freqMap[key],_ = strconv.Atoi(val)
	}
	return freqMap
}

func WriteMatrixToFile(matrix []float64, sampleNames []string, filename string) {
	res := ""
	for _, sample := range sampleNames {
		res = res + sample + "\t"
	}
	res = res + "\n"
	for _, evenness := range matrix {
		res = res + strconv.FormatFloat(evenness, 'f', 3, 64) + "\t"
	}
	ioutil.WriteFile(filename, []byte(res), 0644)
}
